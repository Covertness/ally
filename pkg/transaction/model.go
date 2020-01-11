package transaction

import (
	"time"

	"github.com/Covertness/ally/pkg/address"
	"github.com/cockroachdb/apd"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// Transaction blockchain transaction
type Transaction struct {
	ID        uint `gorm:"primary_key"`
	GroupID   uint `gorm:"index"`
	AddressID uint `gorm:"index:"`
	Address   address.Address
	Status    string `gorm:"index;type:varchar(32);not null"`
	Nonce     string
	Hash      string       `gorm:"unique_index"`
	To        string       `gorm:"index:"`
	Amount    *apd.Decimal `gorm:"type:numeric(38,18);default:0.0;not null"`
	Meta      postgres.Jsonb
	CreatedAt time.Time `gorm:"index;not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

// Transaction status
const (
	StatusInit    = "INIT"
	StatusHolding = "HOLDING" // an error type
	StatusPending = "PENDING"
	StatusDone    = "DONE"
	StatusFailed  = "FAILED"
)
