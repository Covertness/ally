package market

// Market market item
type Market struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"unique_index"`
	Provider    string
	Description string
	Enable      bool `gorm:"default:true;not null"`
}

// Market provider
const (
	ProviderFTX      = "FTX"
	ProviderCoinBase = "Coinbase"
)
