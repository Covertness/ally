package address

import (
	"time"

	"github.com/cockroachdb/apd"
)

// Address blockchain address
type Address struct {
	ID        uint `gorm:"primary_key"`
	AccountID uint `gorm:"index;not null"`

	Address    string       `gorm:"unique_index"`
	Debt       *apd.Decimal `gorm:"type:numeric(38,18);default:0.0;not null"`
	FreezeDebt *apd.Decimal `gorm:"type:numeric(38,18);default:0.0;not null"`
	CreatedAt  time.Time    `gorm:"not null"`
	UpdatedAt  time.Time    `gorm:"not null"`
}
