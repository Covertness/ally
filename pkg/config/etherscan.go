package config

// GetEtherscanHost get the Etherscan host
func GetEtherscanHost() string {
	return getenv("ETHERSCAN_HOST", "api-kovan.etherscan.io")
}

// GetEtherscanAPIKey get the Etherscan api key
func GetEtherscanAPIKey() string {
	return getenv("ETHERSCAN_APIKEY", "")
}
