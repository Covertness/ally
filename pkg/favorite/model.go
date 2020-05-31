package favorite

import "time"

// Favorite favorite item
type Favorite struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt  time.Time    `gorm:"not null"`
	UpdatedAt  time.Time    `gorm:"not null"`

	AccountID uint `gorm:"index;unique_index:uix_favorites_account_item_type_id;not null"`

	ItemType string `gorm:"unique_index:uix_favorites_account_item_type_id;not null"`
	ItemID uint `gorm:"unique_index:uix_favorites_account_item_type_id;not null"`
}

// Favorite item type
const (
	ItemTypeMarket      = "market"
	ItemTypeTimeline = "timeline"
)
