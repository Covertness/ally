package account

import (
	"time"

	"github.com/Covertness/ally/pkg/address"

	"github.com/cockroachdb/apd"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// Account user's account associated with app
type Account struct {
	ID            uint         `gorm:"primary_key"`
	Name          string       `gorm:"unique_index"`
	Balance       *apd.Decimal `gorm:"type:numeric(38,18);default:0.0;not null"`
	FreezeBalance *apd.Decimal `gorm:"type:numeric(38,18);default:0.0;not null"`
	Meta          postgres.Jsonb
	CreatedAt     time.Time `gorm:"not null"`
	UpdatedAt     time.Time `gorm:"not null"`

	Addresses []address.Address
}
