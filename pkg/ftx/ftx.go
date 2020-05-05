package ftx

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/Covertness/ally/pkg/config"
	"github.com/imroc/req"
	"time"
)

var ins *FTX

// Init create the FTX connection
func Init() *FTX {
	ins = New(config.GetFTXKey(), config.GetFTXSecret())
	return ins
}

// GetInstance get the FTX instance
func GetInstance() *FTX {
	return ins
}

// New create a new instance
func New(key, secret string) *FTX {
	return &FTX{
		Endpoint:          "https://ftx.com/api",
		Key: key,
		Secret:          secret,
	}
}

func (f *FTX) get(path string) (*req.Resp, error) {
	url := fmt.Sprintf("%s%s",
		f.Endpoint, path,
	)

	return req.Get(url, f.authHeader("GET", path))
}

func (f *FTX) authHeader(method, path string) req.Header {
	ts := time.Now().Unix()
	signaturePayload := fmt.Sprintf("%d%s%s", ts, method, path)
	h := hmac.New(sha256.New, []byte(f.Secret))
	h.Write([]byte(signaturePayload))
	signature := hex.EncodeToString(h.Sum(nil))

	return req.Header{
		"FTX-KEY":        f.Key,
		"FTX-SIGN": signature,
		"FTX-TS": fmt.Sprint(ts),
	}
}

// FTX instance connect to https://ftx.com
type FTX struct {
	Endpoint          string
	Key          string
	Secret          string
}
