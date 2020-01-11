package config

// GetTokenContract get the token contract address
func GetTokenContract() string {
	return getenv("TOKEN_CONTRACT", "0xFab46E002BbF0b4509813474841E0716E6730136")
}
