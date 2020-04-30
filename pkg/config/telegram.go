package config

// GetTelegramToken get the Telegram Token
func GetTelegramToken() string {
	return getenv("TELEGRAM_TOKEN", "")
}
