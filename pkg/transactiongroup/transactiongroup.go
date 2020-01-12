package transactiongroup

import (
	"errors"

	"github.com/Covertness/ally/pkg/transaction"

	"github.com/Covertness/ally/pkg/address"
	"github.com/Covertness/ally/pkg/config"
	"github.com/Covertness/ally/pkg/storage"

	"github.com/cockroachdb/apd"
)

// CreateAdminWithdrawGroup create withdraw transactions group
func CreateAdminWithdrawGroup(amount *apd.Decimal, to string) (*TransactionGroup, error) {
	tx := storage.GetDB().Begin()

	var addrs []*address.Address
	var err error
	if config.GetDBDialect() == "sqlite3" {
		err = tx.Where("debt > 0").Find(&addrs).Error
	} else {
		err = tx.Set("gorm:query_option", "FOR UPDATE").Where("debt > 0").Find(&addrs).Error
	}
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	newGroup := &TransactionGroup{
		Type: TypeAdminWithdraw,
	}
	err = tx.Create(newGroup).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	totalWithdraw, _, _ := apd.NewFromString(amount.String())
	for _, addr := range addrs {
		var withdraw *apd.Decimal
		if addr.Debt.Cmp(totalWithdraw) >= 0 {
			withdraw, _, _ = apd.NewFromString(totalWithdraw.String())
		} else {
			withdraw, _, _ = apd.NewFromString(addr.Debt.String())
		}

		newTransaction := &transaction.Transaction{
			GroupID:   newGroup.ID,
			AddressID: addr.ID,
			Status:    transaction.StatusInit,
			To:        to,
			Amount:    withdraw,
		}
		err := tx.Create(newTransaction).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		condition, err := apd.BaseContext.Sub(addr.Debt, addr.Debt, withdraw)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		if condition.Any() {
			tx.Rollback()
			return nil, errors.New(condition.String())
		}

		condition, err = apd.BaseContext.Add(addr.FreezeDebt, addr.FreezeDebt, withdraw)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		if condition.Any() {
			tx.Rollback()
			return nil, errors.New(condition.String())
		}

		err = tx.Save(&addr).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		condition, err = apd.BaseContext.Sub(totalWithdraw, totalWithdraw, withdraw)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		if condition.Any() {
			tx.Rollback()
			return nil, errors.New(condition.String())
		}

		if totalWithdraw.IsZero() {
			break
		}
	}

	if !totalWithdraw.IsZero() {
		tx.Rollback()
		return nil, errors.New("not enough balance")
	}

	tx.Commit()
	return newGroup, nil
}

// GetByID get transactions by group id
func GetByID(id uint) (*TransactionGroup, error) {
	var model TransactionGroup
	err := storage.GetDB().Preload("Transactions.Address").First(&model, id).Error
	if err != nil {
		return nil, err
	}

	return &model, err
}
