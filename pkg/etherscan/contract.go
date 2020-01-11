package etherscan

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/Covertness/ally/contracts"
	"github.com/Covertness/ally/pkg/ethclient"
)

// GetTokenInfo get token info
func (e *Etherscan) GetTokenInfo() (*contracts.ERC20, error) {
	eth := ethclient.GetInstance()
	return eth.GetERC20Token(e.ContractAddress)
}

// GetTokenDecimals get token decimals
func (e *Etherscan) GetTokenDecimals() (uint8, error) {
	if e.TokenDecimals > 0 {
		return e.TokenDecimals, nil
	}

	token, err := e.GetTokenInfo()
	if err != nil {
		return 0, err
	}

	e.TokenDecimals, err = token.Decimals(&bind.CallOpts{})
	return e.TokenDecimals, err
}
