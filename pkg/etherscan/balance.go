package etherscan

import (
	"errors"
	"fmt"

	"github.com/cockroachdb/apd"
	"github.com/imroc/req"
)

// GetBalance get ether balance for a specified address
func (e *Etherscan) GetBalance(address string) (*apd.Decimal, error) {
	url := fmt.Sprintf("https://%s/api?module=account&action=balance&address=%s&tag=latest&apikey=%s",
		e.Domain, address, e.APIKey,
	)
	res, err := req.Get(url)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	res.ToJSON(&result)
	if result["status"] != "1" {
		return nil, errors.New(result["result"].(string))
	}

	b, _, _ := apd.NewFromString(result["result"].(string))
	return b, nil
}

// GetTokenBalance get token balance for a specified address
func (e *Etherscan) GetTokenBalance(address string) (*apd.Decimal, error) {
	url := fmt.Sprintf("https://%s/api?module=account&action=tokenbalance&contractaddress=%s&address=%s&tag=latest&apikey=%s",
		e.Domain, e.ContractAddress, address, e.APIKey,
	)
	res, err := req.Get(url)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	res.ToJSON(&result)
	if result["status"] != "1" {
		return nil, errors.New(result["result"].(string))
	}

	b, _, _ := apd.NewFromString(result["result"].(string))
	return b, nil
}
