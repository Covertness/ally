package order

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/Covertness/ally/pkg/account"
	"github.com/Covertness/ally/pkg/address"
	"github.com/Covertness/ally/pkg/ethclient"
	"github.com/Covertness/ally/pkg/item"
	"github.com/Covertness/ally/pkg/storage"

	"github.com/cockroachdb/apd"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// Create create object in database
func Create(sequence uint, mAccount *account.Account, mItem *item.Item) (*Order, error) {
	tx := storage.GetDB().Begin()

	itemByte, _ := json.Marshal(mItem)

	newOrder := &Order{
		Sequence:     sequence,
		AccountID:    mAccount.ID,
		ItemID:       mItem.ID,
		ItemSnapshot: postgres.Jsonb{RawMessage: itemByte},
		Status:       StatusInit,
	}
	err := tx.Create(newOrder).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return newOrder, nil
}

// GetByStatuses get orders by status
func GetByStatuses(statuses []string) (orders []*Order, err error) {
	err = storage.GetDB().Where("status IN (?)", statuses).Find(&orders).Error
	return orders, err
}

// GetByID get order by id
func GetByID(id uint) (*Order, error) {
	var model Order
	err := storage.GetDB().Preload("Address").First(&model, id).Error
	if err != nil {
		return nil, err
	}

	return &model, err
}

// RequestAddress request address from eth contract
func (o *Order) RequestAddress() error {
	tx := storage.GetDB().Begin()

	var latestOrder Order
	err := tx.Set("gorm:query_option", "FOR UPDATE").First(&latestOrder, o.ID).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if latestOrder.Status == StatusInit {
		addr, err := o.Account.RequestAddress(tx)
		if err != nil {
			tx.Rollback()
			return err
		}

		err = tx.Model(&latestOrder).Updates(Order{
			AddressID: addr.ID,
			Status:    StatusWaitingAddress,
		}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if latestOrder.Status == StatusWaitingAddress {
		success, err := o.fetchAddress(tx)
		if err != nil {
			tx.Rollback()
			return err
		}

		if success {
			err = tx.Model(&latestOrder).Updates(Order{
				Status: StatusDepositable,
			}).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	tx.Commit()
	return nil
}

// ApplyDeposit add account balance
func (o *Order) ApplyDeposit(amount *apd.Decimal) error {
	tx := storage.GetDB().Begin()

	var addr address.Address
	err := tx.Set("gorm:query_option", "FOR UPDATE").First(&addr, o.AddressID).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	var acc account.Account
	err = tx.Set("gorm:query_option", "FOR UPDATE").First(&acc, addr.AccountID).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	var latestOrder Order
	err = tx.Set("gorm:query_option", "FOR UPDATE").First(&latestOrder, o.ID).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if latestOrder.Status == StatusDone {
		tx.Rollback()
		return nil
	}

	var balance apd.Decimal
	condition, err := apd.BaseContext.Sub(&balance, amount, addr.Debt)
	if err != nil {
		tx.Rollback()
		return err
	}
	if condition.Any() {
		tx.Rollback()
		return errors.New(condition.String())
	}

	condition, err = apd.BaseContext.Add(addr.Debt, addr.Debt, &balance)
	if err != nil {
		tx.Rollback()
		return err
	}
	if condition.Any() {
		tx.Rollback()
		return errors.New(condition.String())
	}

	condition, err = apd.BaseContext.Add(acc.Balance, acc.Balance, &balance)
	if err != nil {
		tx.Rollback()
		return err
	}
	if condition.Any() {
		tx.Rollback()
		return errors.New(condition.String())
	}

	price, err := o.GetItemPrice()
	if err != nil {
		tx.Rollback()
		return err
	}

	condition, err = apd.BaseContext.Sub(acc.Balance, acc.Balance, price)
	if err != nil {
		tx.Rollback()
		return err
	}
	if condition.Any() {
		tx.Rollback()
		return errors.New(condition.String())
	}

	err = tx.Save(&acc).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Model(o).Updates(Order{
		Status: StatusDone,
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Save(&addr).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// GetItemPrice the price from snapshot
func (o *Order) GetItemPrice() (*apd.Decimal, error) {
	var mItem item.Item
	err := json.Unmarshal(o.ItemSnapshot.RawMessage, &mItem)
	if err != nil {
		return nil, err
	}
	return mItem.Price, nil
}

func (o *Order) fetchAddress(tx *gorm.DB) (bool, error) {
	eth := ethclient.GetInstance()
	factory, err := eth.GetAddressFactory()
	if err != nil {
		return false, err
	}

	id := new(big.Int).SetUint64(uint64(o.Address.ID))

	addr, err := factory.ReceiversMap(nil, id)
	if err != nil {
		return false, err
	}

	if addr.Hash().Big().Sign() == 0 {
		return false, nil
	}

	err = tx.Model(&o.Address).Updates(address.Address{
		Address: addr.Hex(),
	}).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
