package config

// GetFTXKey get the FTX Key
func GetFTXKey() string {
	return getenv("FTX_KEY", "")
}

// GetFTXSecret get the FTX Secret
func GetFTXSecret() string {
	return getenv("FTX_SECRET", "")
}
