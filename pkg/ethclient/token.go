package ethclient

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/Covertness/ally/contracts"
)

// GetERC20Token get ERC20 token info from token address
func (c *Client) GetERC20Token(tokenAddress string) (*contracts.ERC20, error) {
	if c.myERC20Token != nil {
		return c.myERC20Token, nil
	}

	address := common.HexToAddress(tokenAddress)
	token, err := contracts.NewERC20(address, c.Client)

	c.myERC20Token = token
	return c.myERC20Token, err
}
