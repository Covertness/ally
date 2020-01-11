package ethclient

import (
	"os"

	"github.com/ethereum/go-ethereum/common"

	"github.com/Covertness/ally/contracts"
)

// GetAddressFactory get contract instance AddressFactory
func (c *Client) GetAddressFactory() (*contracts.AddressFactory, error) {
	if c.myAddressFactory != nil {
		return c.myAddressFactory, nil
	}

	factory, err := contracts.NewAddressFactory(
		common.HexToAddress(os.Getenv("CONTRACT_ADDRESS_FACTORY_ADDRESS")),
		c.Client,
	)
	c.myAddressFactory = factory
	return c.myAddressFactory, err
}
