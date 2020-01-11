package order

import (
	"time"

	"github.com/Covertness/ally/pkg/account"

	"github.com/Covertness/ally/pkg/address"
	"github.com/Covertness/ally/pkg/item"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// Order order a product
type Order struct {
	ID           uint `gorm:"primary_key"`
	Sequence     uint `gorm:"unique_index;not null"`
	AccountID    uint `gorm:"index;not null"`
	Account      account.Account
	AddressID    uint `gorm:"index"`
	Address      address.Address
	ItemID       uint `gorm:"index;not null"`
	Item         item.Item
	ItemSnapshot postgres.Jsonb
	Status       string    `gorm:"index;type:varchar(32);not null"`
	CreatedAt    time.Time `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"index;not null"`
}

// Order status
const (
	StatusInit           = "INIT"
	StatusWaitingAddress = "WAITING_ADDRESS"
	StatusDepositable    = "DEPOSITABLE"
	StatusDepositing     = "DEPOSITING"
	StatusDone           = "DONE"
	StatusTimeout        = "TIMEOUT"
)
