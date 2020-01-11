package item

import (
	"time"

	"github.com/cockroachdb/apd"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// Item represent a product
type Item struct {
	ID         uint         `gorm:"primary_key"`
	ExternalID string       `gorm:"unique_index"`
	Price      *apd.Decimal `gorm:"type:numeric(38,18);default:0.0;not null"`
	Meta       postgres.Jsonb
	CreatedAt  time.Time `gorm:"not null"`
	UpdatedAt  time.Time `gorm:"not null"`
}
