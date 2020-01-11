package etherscan

import (
	"github.com/Covertness/ally/pkg/config"
)

var ins *Etherscan

// Init create the Etherscan connection
func Init() *Etherscan {
	ins = New(config.GetEtherscanHost(), config.GetTokenContract(), config.GetEtherscanAPIKey())
	return ins
}

// GetInstance get the Etherscan instance
func GetInstance() *Etherscan {
	return ins
}

// New create a new instance
func New(domain, contractAddress, apiKey string) *Etherscan {
	return &Etherscan{
		Domain:          domain,
		ContractAddress: contractAddress,
		APIKey:          apiKey,
	}
}

// Etherscan instance connect to https://etherscan.io
type Etherscan struct {
	Domain          string
	ContractAddress string
	APIKey          string
	TokenDecimals   uint8
}
