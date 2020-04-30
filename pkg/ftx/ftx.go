package ftx

import "github.com/Covertness/ally/pkg/config"

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

// FTX instance connect to https://ftx.com
type FTX struct {
	Endpoint          string
	Key          string
	Secret          string
}
