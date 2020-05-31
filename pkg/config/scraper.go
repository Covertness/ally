package config

// GetScraperEndpoint get the Scraper Endpoint
func GetScraperEndpoint() string {
	return getenv("SCRAPER_ENDPOINT", "http://127.0.0.1:3000")
}
