package interchaintest_test

import (
	"testing"

	"github.com/strangelove-ventures/interchaintest/v3/ibc"
	integration "github.com/strangelove-ventures/noble/interchaintest"
)

func TestGrand1ChainUpgrade(t *testing.T) {
	repo, version := integration.GetDockerImageInfo()

	const (
		grand1ChainID = "noble-1"
		numVals       = 4
		numFullNodes  = 0
	)

	var grand1Genesis = ibc.DockerImage{
		Repository: "ghcr.io/strangelove-ventures/noble",
		Version:    "v0.3.0",
		UidGid:     containerUidGid,
	}

	var grand1Upgrades = []chainUpgrade{
		{
			upgradeName: "neon",
			image: ibc.DockerImage{
				Repository: "ghcr.io/strangelove-ventures/noble",
				Version:    "v0.4.2",
				UidGid:     containerUidGid,
			},
		},
		{
			upgradeName: "radon",
			image: ibc.DockerImage{
				Repository: "ghcr.io/strangelove-ventures/noble",
				// testnet actually upgraded to v0.5.0, but that required a hack to fix the consensus min fee. v0.5.1 fixes that.
				Version: "v0.5.1",
				UidGid:  containerUidGid,
			},
			postUpgrade: testPostRadonUpgrade,
		},
		{
			// post radon patch upgrade (will be applied as rolling upgrade due to lack of upgradeName)
			image: ibc.DockerImage{
				Repository: repo,
				Version:    version,
				UidGid:     containerUidGid,
			},
		},
	}

	testNobleChainUpgrade(t, grand1ChainID, grand1Genesis, denomMetadataDrachma, numVals, numFullNodes, grand1Upgrades)
}
