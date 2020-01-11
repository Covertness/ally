package etherscan

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/imroc/req"
)

// GetTransaction get the transaction by hash
func (e *Etherscan) GetTransaction(hash string) (*Transaction, error) {
	url := fmt.Sprintf("https://%s/api?module=proxy&action=eth_getTransactionByHash&txhash=%s&apikey=%s",
		e.Domain, hash, e.APIKey,
	)
	res, err := req.Get(url)
	if err != nil {
		return nil, err
	}

	var result transactionResponse
	err = res.ToJSON(&result)
	if err != nil {
		return nil, err
	}
	if result.Result.Hash != hash {
		return nil, errors.New("not found the transaction")
	}

	value, _ := new(big.Int).SetString(result.Result.Value, 0)
	gas, _ := new(big.Int).SetString(result.Result.Gas, 0)
	gasPrice, _ := new(big.Int).SetString(result.Result.GasPrice, 0)
	blockNumber, _ := new(big.Int).SetString(result.Result.BlockNumber, 0)

	return &Transaction{
		Hash:        result.Result.Hash,
		From:        result.Result.From,
		To:          result.Result.To,
		Value:       value,
		Gas:         gas,
		GasPrice:    gasPrice,
		BlockNumber: blockNumber,
	}, nil
}

type transactionResponse struct {
	Web3Response

	Result struct {
		BlockNumber string
		Hash        string
		From        string
		To          string
		Value       string
		Gas         string
		GasPrice    string
	}
}

// Transaction ETH transaction
type Transaction struct {
	Hash        string
	From        string
	To          string
	Value       *big.Int
	Gas         *big.Int
	GasPrice    *big.Int
	BlockNumber *big.Int
}
