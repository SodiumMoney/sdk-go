package sodium

import "sync"

type networkConfig struct {
	Label     string
	ProgramID string
}

var (
	networks = map[Network]networkConfig{
		Devnet:  {Label: "Devnet", ProgramID: "4qduM91VaZj9W9hRypygozjN7ZCeLPPe5veM2RhB7qgD"},
		Mainnet: {Label: "Mainnet-beta", ProgramID: "FQr13NtzFwLp8ZZwR8SWfjEGC4F4MBrJDrcSZuEsv3Gx"},
	}
	DefaultNetwork = Mainnet
	netMu          sync.RWMutex
)

func GetProgramID(network Network) string {
	netMu.RLock()
	defer netMu.RUnlock()
	return networks[network].ProgramID
}

// SetMainnetProgramID updates the mainnet program id after deploy.
func SetMainnetProgramID(id string) {
	netMu.Lock()
	defer netMu.Unlock()
	cfg := networks[Mainnet]
	cfg.ProgramID = id
	networks[Mainnet] = cfg
}
