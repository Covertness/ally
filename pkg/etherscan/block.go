package etherscan

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/imroc/req"
)

// GetBlockNumber get the current latest block number
func (e *Etherscan) GetBlockNumber() (*big.Int, error) {
	url := fmt.Sprintf("https://%s/api?module=proxy&action=eth_blockNumber&apikey=%s",
		e.Domain, e.APIKey,
	)
	res, err := req.Get(url)
	if err != nil {
		return nil, err
	}

	var result blockNumberResponse
	err = res.ToJSON(&result)
	if err != nil {
		return nil, err
	}
	if len(result.Result) == 0 {
		return nil, errors.New("get block number failed")
	}

	blockNumber, _ := new(big.Int).SetString(result.Result, 0)

	return blockNumber, nil
}

type blockNumberResponse struct {
	Web3Response

	Result string
}
