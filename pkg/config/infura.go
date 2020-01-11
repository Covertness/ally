package config

// GetInfuraEndpoint get the Infura endpoint
func GetInfuraEndpoint() string {
	return getenv("INFURA_ENDPOINT", "kovan")
}

// GetInfuraID get the Infura id
func GetInfuraID() string {
	return getenv("INFURA_ID", "")
}
