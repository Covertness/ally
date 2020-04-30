package ftx

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/imroc/req"
	"time"
)

func (f *FTX) GetMarket(name string) (*Market, error) {
	path := fmt.Sprintf("/markets/%s", name)
	url := fmt.Sprintf("%s%s",
		f.Endpoint, path,
	)

	ts := time.Now().Unix()
	signaturePayload := fmt.Sprintf("%d%s%s", ts, "GET", path)
	h := hmac.New(sha256.New, []byte(f.Secret))
	h.Write([]byte(signaturePayload))
	signature := hex.EncodeToString(h.Sum(nil))

	authHeader := req.Header{
		"FTX-KEY":        f.Key,
		"FTX-SIGN": signature,
		"FTX-TS": fmt.Sprint(ts),
	}
	r, err := req.Get(url, authHeader)
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

type Market struct {
	Name string
	Type string
	Ask float64
	Bid float64
	Last float64
}
