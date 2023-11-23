package interchaintest_test

import (
	"context"
	"encoding/json"
	"strconv"
	"testing"
	"time"

	ibcclienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	providerclient "github.com/cosmos/interchain-security/v2/x/ccv/provider/client"
	"github.com/icza/dyno"
	"github.com/strangelove-ventures/interchaintest/v4"
	"github.com/strangelove-ventures/interchaintest/v4/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v4/ibc"
	"github.com/strangelove-ventures/interchaintest/v4/relayer/hermes"
	"github.com/strangelove-ventures/interchaintest/v4/testreporter"
	"github.com/strangelove-ventures/interchaintest/v4/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func modifyGaiaGenesis(_ ibc.ChainConfig, bz []byte) ([]byte, error) {
	genesis := make(map[string]interface{})

	if err := json.Unmarshal(bz, &genesis); err != nil {
		return nil, err
	}

	err := dyno.Set(genesis, (15 * time.Second).String(), "app_state", "gov", "voting_params", "voting_period")
	if err != nil {
		return nil, err
	}

	return json.Marshal(genesis)
}

func TestICS_StandaloneToConsumerMigration(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	logger := zaptest.NewLogger(t)
	reporter := testreporter.NewNopReporter()
	execReporter := reporter.RelayerExecReporter(t)
	client, network := interchaintest.DockerSetup(t)

	var wrapper genesisWrapper

	spec := nobleChainSpec(ctx, &wrapper, "grand-rehearsal-1", 1, 0, false, false, false, false)
	spec.Images = []ibc.DockerImage{ghcrImage("v4.0.1")}

	factory := interchaintest.NewBuiltinChainFactory(logger, []*interchaintest.ChainSpec{
		spec,
		{
			Name:    "gaia",
			Version: "latest",
			ChainConfig: ibc.ChainConfig{
				ModifyGenesis: modifyGaiaGenesis,
			},
		},
	})

	chains, err := factory.Chains(t.Name())
	require.NoError(t, err)

	noble := chains[0].(*cosmos.CosmosChain)
	gaia := chains[1].(*cosmos.CosmosChain)
	wrapper.chain = noble

	rly := interchaintest.NewBuiltinRelayerFactory(
		ibc.Hermes,
		logger,
	).Build(t, client, network).(*hermes.Relayer)

	interchain := interchaintest.NewInterchain().
		AddChain(noble).
		AddChain(gaia).
		AddRelayer(rly, "rly").
		AddLink(interchaintest.InterchainLink{
			Chain1:  noble,
			Chain2:  gaia,
			Relayer: rly,
			Path:    "transfer",
		})

	require.NoError(t, interchain.Build(ctx, execReporter, interchaintest.InterchainBuildOptions{
		TestName:  t.Name(),
		Client:    client,
		NetworkID: network,
	}))

	t.Cleanup(func() {
		_ = interchain.Close()
	})

	// Register all Gaia validators as Noble full nodes.
	for _, val := range gaia.Validators {
		content, err := val.PrivValFileContent(ctx)
		require.NoError(t, err)

		require.NoError(t, noble.AddFullNodes(ctx, nil, 1))
		require.NoError(t, noble.FullNodes[noble.NumFullNodes-1].OverwritePrivValFile(ctx, content))
	}

	// Register the Krypton upgrade on Noble.
	upgradeHeight := 50

	_, err = noble.Validators[0].ExecTx(ctx, wrapper.paramAuthority.KeyName(), "upgrade", "software-upgrade", "krypton", "--upgrade-height", strconv.Itoa(upgradeHeight))
	require.NoError(t, err)

	current, err := noble.Height(ctx)
	require.NoError(t, err)

	timeoutCtx, timeoutCtxCancel := context.WithTimeout(ctx, 15*time.Second)
	defer timeoutCtxCancel()

	_ = testutil.WaitForBlocks(timeoutCtx, (upgradeHeight-int(current))+1, noble)

	current, err = noble.Height(ctx)
	require.NoError(t, err)
	require.Equal(t, current, uint64(upgradeHeight))

	// Register Noble as a consumer chain.
	proposers := interchaintest.GetAndFundTestUsers(t, ctx, "proposer", 100_000_000, gaia)
	proposer := proposers[0]

	channels, err := rly.GetChannels(ctx, execReporter, gaia.Config().ChainID)
	require.NoError(t, err)
	require.Equal(t, len(channels), 1)

	proposal, err := gaia.ConsumerAdditionProposal(ctx, proposer.KeyName(), providerclient.ConsumerAdditionProposalJSON{
		Title:                             "Migrate Noble to ICS",
		Description:                       "This proposal aims to migrate Noble from a standalone POA chain to ICS.",
		ChainId:                           noble.Config().ChainID,
		InitialHeight:                     ibcclienttypes.NewHeight(1, uint64(upgradeHeight+1)),
		GenesisHash:                       []byte("genesis"),
		BinaryHash:                        []byte("binary"),
		SpawnTime:                         time.Now(),
		ConsumerRedistributionFraction:    "0.85",
		BlocksPerDistributionTransmission: 1000,
		DistributionTransmissionChannel:   channels[0].ChannelID,
		HistoricalEntries:                 10_000,
		CcvTimeoutPeriod:                  time.Hour * 24 * 7 * 4, // 2419200s
		TransferTimeoutPeriod:             time.Hour,              // 3600s
		UnbondingPeriod:                   time.Hour * 24 * 7 * 3, // 1814400s
		Deposit:                           "10000000uatom",
	})
	require.NoError(t, err)

	require.NoError(t, gaia.VoteOnProposalAllValidators(ctx, proposal.ProposalID, cosmos.ProposalVoteYes))

	current, _ = gaia.Height(ctx)
	_, err = cosmos.PollForProposalStatus(ctx, gaia, current, current+100, proposal.ProposalID, cosmos.ProposalStatusPassed)
	require.NoError(t, err)

	// Wait for 1 block to be produced on Gaia.
	// This is so we can query the consumer genesis information.
	require.NoError(t, testutil.WaitForBlocks(ctx, 1, gaia))

	// Add consumer genesis information to all Noble nodes.
	ccv, _, err := gaia.FullNodes[0].ExecQuery(ctx, "provider", "consumer-genesis", noble.Config().ChainID)
	require.NoError(t, err)

	for _, node := range noble.Nodes() {
		require.NoError(t, node.WriteFile(ctx, ccv, "config/ccv.json"))
	}

	// Execute the Krypton upgrade on Noble.
	require.NoError(t, noble.StopAllNodes(ctx))

	container := nobleImageInfo[0]
	noble.UpgradeVersion(ctx, client, container.Repository, container.Version)

	require.NoError(t, noble.StartAllNodes(ctx))

	// Create the Replicated Security channel using the newly created clients.
	var res ibc.RelayerExecResult
	nobleClientID, gaiaClientID := "07-tendermint-1", "07-tendermint-1"

	res = rly.Exec(ctx, execReporter, []string{"hermes", "--json", "create", "connection", "--a-chain", noble.Config().ChainID, "--a-client", nobleClientID, "--b-client", gaiaClientID}, nil)
	require.NoError(t, res.Err)

	nobleConnectionID, _, err := hermes.GetConnectionIDsFromStdout(res.Stdout)
	require.NoError(t, err)

	res = rly.Exec(ctx, execReporter, []string{"hermes", "create", "channel", "--a-chain", noble.Config().ChainID, "--a-connection", nobleConnectionID, "--a-port", "consumer", "--b-port", "provider", "--order", "ORDER_ORDERED"}, nil)
	require.NoError(t, res.Err)

	// Execute a delegation on Gaia to generate a VSC packet.
	delegators := interchaintest.GetAndFundTestUsers(t, ctx, "delegator", 1000000000000, gaia)
	delegator := delegators[0]

	validator, err := gaia.Validators[0].KeyBech32(ctx, "validator", "val")
	require.NoError(t, err)

	_, err = gaia.FullNodes[0].ExecTx(ctx, delegator.KeyName(), "staking", "delegate", validator, "999999000000uatom")
	require.NoError(t, err)

	// Clear packets on the Replicated Security channel.
	res = rly.Exec(ctx, execReporter, []string{"hermes", "clear", "packets", "--chain", noble.Config().ChainID, "--port", "consumer", "--channel", "channel-1"}, nil)
	require.NoError(t, res.Err)

	// Wait for 10 blocks to be produced.
	require.NoError(t, testutil.WaitForBlocks(ctx, 10, noble))
}
