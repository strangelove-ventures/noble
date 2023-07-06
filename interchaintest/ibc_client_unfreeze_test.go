package interchaintest_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/strangelove-ventures/interchaintest/v3"
	"github.com/strangelove-ventures/interchaintest/v3/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v3/ibc"
	"github.com/strangelove-ventures/interchaintest/v3/testreporter"
	"github.com/strangelove-ventures/interchaintest/v3/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

// run `make local-image`to rebuild updated binary before running test
func TestClientUnfreeze(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Parallel()

	ctx := context.Background()

	rep := testreporter.NewNopReporter()
	eRep := rep.RelayerExecReporter(t)

	client, network := interchaintest.DockerSetup(t)

	var (
		noble                *cosmos.CosmosChain
		roles                NobleRoles
		roles2               NobleRoles
		paramauthorityWallet ibc.Wallet
	)

	chainCfg := ibc.ChainConfig{
		Type:           "cosmos",
		Name:           "noble",
		ChainID:        "noble-1",
		Bin:            "nobled",
		Denom:          "token",
		Bech32Prefix:   "noble",
		CoinType:       "118",
		GasPrices:      "0.0token",
		GasAdjustment:  1.1,
		TrustingPeriod: "504h",
		NoHostMount:    false,
		Images:         nobleImageInfo,
		EncodingConfig: NobleEncoding(),
		PreGenesis: func(cc ibc.ChainConfig) (err error) {
			val := noble.Validators[0]
			err = createTokenfactoryRoles(ctx, &roles, denomMetadataRupee, val, false)
			if err != nil {
				return err
			}
			err = createTokenfactoryRoles(ctx, &roles2, denomMetadataDrachma, val, false)
			if err != nil {
				return err
			}
			paramauthorityWallet, err = createParamAuthAtGenesis(ctx, val)
			return err
		},
		ModifyGenesis: func(cc ibc.ChainConfig, b []byte) ([]byte, error) {
			g := make(map[string]interface{})
			if err := json.Unmarshal(b, &g); err != nil {
				return nil, fmt.Errorf("failed to unmarshal genesis file: %w", err)
			}
			if err := modifyGenesisTokenfactory(g, "tokenfactory", denomMetadataRupee, &roles, true); err != nil {
				return nil, err
			}
			if err := modifyGenesisTokenfactory(g, "fiat-tokenfactory", denomMetadataDrachma, &roles2, true); err != nil {
				return nil, err
			}
			if err := modifyGenesisParamAuthority(g, paramauthorityWallet.FormattedAddress()); err != nil {
				return nil, err
			}
			if err := modifyGenesisTariffDefaults(g, paramauthorityWallet.FormattedAddress()); err != nil {
				return nil, err
			}
			out, err := json.Marshal(&g)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal genesis bytes to json: %w", err)
			}
			return out, nil
		},
	}

	nv := 1
	nf := 0

	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{
			ChainConfig:   chainCfg,
			NumValidators: &nv,
			NumFullNodes:  &nf,
		},
		{
			Name:          "gaia",
			Version:       "v10.0.2",
			NumValidators: &nv,
			NumFullNodes:  &nf,
		},
	})

	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)

	noble = chains[0].(*cosmos.CosmosChain)
	gaia := chains[1].(*cosmos.CosmosChain)

	r := interchaintest.NewBuiltinRelayerFactory(
		ibc.CosmosRly,
		zaptest.NewLogger(t),
		relayerImage,
	).Build(t, client, network)

	pathName := "noble-gaia"

	ic := interchaintest.NewInterchain().
		AddChain(noble).
		AddChain(gaia).
		AddRelayer(r, "r").
		AddLink(interchaintest.InterchainLink{
			Chain1:  noble,
			Chain2:  gaia,
			Relayer: r,
			Path:    pathName,
			CreateClientOpts: ibc.CreateClientOptions{
				TrustingPeriod: "20s",
			},
		})

	require.NoError(t, ic.Build(ctx, eRep, interchaintest.InterchainBuildOptions{
		TestName:  t.Name(),
		Client:    client,
		NetworkID: network,

		// TODO comment out for CI
		BlockDatabaseFile: interchaintest.DefaultBlockDatabaseFilepath(),

		SkipPathCreation: true,
	}))
	t.Cleanup(func() {
		_ = ic.Close()
	})

	const userFunds = int64(10_000_000_000)
	users := interchaintest.GetAndFundTestUsers(t, ctx, t.Name(), userFunds, noble, gaia)

	nobleChainID := noble.Config().ChainID

	nobleClients, err := r.GetClients(ctx, eRep, nobleChainID)
	require.NoError(t, err)
	require.Len(t, nobleClients, 1)

	nobleClient := nobleClients[0]

	nobleChannels, err := r.GetChannels(ctx, eRep, nobleChainID)
	require.NoError(t, err)
	require.Len(t, nobleChannels, 1)
	nobleChannel := nobleChannels[0]

	err = testutil.WaitForBlocks(ctx, 20, noble)
	require.NoError(t, err)

	// client should now be expired, no relayer was running to update the clients during the 20s trusting period.

	_, err = noble.SendIBCTransfer(ctx, nobleChannel.ChannelID, users[0].KeyName(), ibc.WalletAmount{
		Address: users[1].FormattedAddress(),
		Amount:  1000000,
		Denom:   noble.Config().Denom,
	}, ibc.TransferOptions{})

	require.Error(t, err)
	require.ErrorContains(t, err, "status Expired: client is not active")

	// create new client
	require.NoError(t, r.CreateClients(ctx, eRep, pathName, ibc.CreateClientOptions{TrustingPeriod: "1h", Override: true}))

	nobleClients, err = r.GetClients(ctx, eRep, nobleChainID)
	require.NoError(t, err)
	require.Len(t, nobleClients, 2)

	newNobleClient := nobleClients[1]

	// substitute new client for old client
	_, err = noble.Validators[0].ExecTx(ctx, paramauthorityWallet.KeyName(), "upgrade", "update-client", nobleClient.ClientID, newNobleClient.ClientID)
	require.NoError(t, err)

	// send a packet on the same channel but new client, should succeed.
	_, err = noble.SendIBCTransfer(ctx, nobleChannel.ChannelID, users[0].KeyName(), ibc.WalletAmount{
		Address: users[1].FormattedAddress(),
		Amount:  1000000,
		Denom:   noble.Config().Denom,
	}, ibc.TransferOptions{})
	require.NoError(t, err)

}
