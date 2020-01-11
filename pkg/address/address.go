package address

import (
	"github.com/Covertness/ally/pkg/storage"
	"github.com/cockroachdb/apd"
)

// GetTotalDebts get all debts belong to addresses
func GetTotalDebts() (*apd.Decimal, error) {
	var debt debtResult
	err := storage.GetDB().Table(storage.GetDB().NewScope(Address{}).TableName()).
		Select("sum(debt) as amount").Scan(&debt).Error
	return debt.Amount, err
}

type debtResult struct {
	Amount *apd.Decimal
}
