package transactiongroup

import (
	"github.com/Covertness/ally/pkg/transaction"
)

// TransactionGroup aggregate transactions
type TransactionGroup struct {
	ID uint `gorm:"primary_key"`

	Type string `gorm:"index;type:varchar(32)"`

	Transactions []*transaction.Transaction `gorm:"foreignkey:GroupID"`
}

// TransactionGroup type
const (
	// address
	TypeAdminWithdraw = "ADMIN_WITHDRAW"
)
