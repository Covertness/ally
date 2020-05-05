package ftx

import (
	"errors"
	"fmt"
)

// GetMarkets get all markets
func (f *FTX) GetMarkets() ([]*Market, error) {
	path := fmt.Sprintf("/markets")
	r, err := f.get(path)
	if err != nil {
		return nil, err
	}

	var resp respMarkets
	err = r.ToJSON(&resp)
	if err != nil {
		return nil, err
	}
	if !resp.Success {
		return nil, errors.New(fmt.Sprintf("response failed: %v", resp))
	}

	return resp.Result, nil
}

type respMarkets struct {
	Success bool
	Result []*Market
}

// GetMarket get market by specified name
func (f *FTX) GetMarket(name string) (*Market, error) {
	path := fmt.Sprintf("/markets/%s", name)
	r, err := f.get(path)
	if err != nil {
		return nil, err
	}

	var resp respMarket
	err = r.ToJSON(&resp)
	if err != nil {
		return nil, err
	}
	if !resp.Success {
		return nil, errors.New(fmt.Sprintf("response failed: %v", resp))
	}

	return &resp.Result, nil
}

type respMarket struct {
	Success bool
	Result Market
}

// Market market item
type Market struct {
	Name string
	Type string
	Ask float64
	Bid float64
	Last float64
}
