package coinbase

import (
	"fmt"

	"github.com/imroc/req"
)

var ins *CoinBase

// Init create the CoinBase connection
func Init() *CoinBase {
	ins = New()
	return ins
}

// GetInstance get the CoinBase instance
func GetInstance() *CoinBase {
	return ins
}

// New create a new instance
func New() *CoinBase {
	return &CoinBase{
		Endpoint: "https://api.coinbase.com",
	}
}

func (c *CoinBase) get(path string) (*req.Resp, error) {
	url := fmt.Sprintf("%s%s",
		c.Endpoint, path,
	)

	return req.Get(url)
}

// CoinBase instance connect to https://www.coinbase.com/
type CoinBase struct {
	Endpoint string
}
