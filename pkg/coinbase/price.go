package coinbase

import "fmt"

// GetPriceSpot get the price of currencyPair
func (c *CoinBase) GetPriceSpot(currencyPair string) (*PriceSpot, error) {
	path := fmt.Sprintf("/v2/prices/%s/spot", currencyPair)
	r, err := c.get(path)
	if err != nil {
		return nil, err
	}

	var resp respPriceSpot
	err = r.ToJSON(&resp)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

type respPriceSpot struct {
	Data PriceSpot
}

// PriceSpot price item
type PriceSpot struct {
	Amount   string
	Currency string
}
