package transaction

import (
	"errors"
	"math/big"

	"github.com/Covertness/ally/pkg/address"

	"github.com/cockroachdb/apd"

	"github.com/ethereum/go-ethereum/common"

	"github.com/Covertness/ally/pkg/ethclient"
	"github.com/Covertness/ally/pkg/etherscan"
	"github.com/Covertness/ally/pkg/storage"
)

// GetByStatuses get orders by status
func GetByStatuses(statuses []string) (transactions []*Transaction, err error) {
	err = storage.GetDB().Where("status IN (?)", statuses).Find(&transactions).Error
	return transactions, err
}

// RequestApply create the transaction on the blockchain
func (t *Transaction) RequestApply() error {
	eth := ethclient.GetInstance()
	factory, err := eth.GetAddressFactory()
	if err != nil {
		return err
	}

	if len(t.Nonce) == 0 {
		nonce, err := eth.GenNonce()
		if err != nil {
			return err
		}

		err = storage.GetDB().Model(t).Updates(Transaction{
			Nonce:  nonce.String(),
			Status: StatusHolding,
		}).Error
		if err != nil {
			return err
		}
	}

	nonce := new(big.Int)
	_, ok := nonce.SetString(t.Nonce, 10)
	if !ok {
		return errors.New("nonce error")
	}

	signer, err := eth.GetSigner(nonce)
	if err != nil {
		return err
	}

	scaner := etherscan.GetInstance()
	decimals, err := scaner.GetTokenDecimals()
	if err != nil {
		return err
	}

	var amount apd.Decimal
	condition, err := apd.BaseContext.Mul(&amount, t.Amount, apd.New(1, int32(decimals)))
	if err != nil {
		return err
	}
	if condition.Any() {
		return errors.New(condition.String())
	}

	id := new(big.Int).SetUint64(uint64(t.Address.ID))
	tokenAddress := common.HexToAddress(scaner.ContractAddress)
	toAddress := common.HexToAddress(t.To)

	newTx, err := factory.SendFundsFromReceiverTo(signer, id, tokenAddress, &amount.Coeff, toAddress)
	if err != nil {
		return err
	}

	return storage.GetDB().Model(t).Updates(Transaction{
		Hash:   newTx.Hash().String(),
		Status: StatusPending,
	}).Error
}

// Apply apply the transaction in local data
func (t *Transaction) Apply() error {
	tx := storage.GetDB().Begin()

	addr := t.Address

	var lockAddr address.Address
	err := tx.Set("gorm:query_option", "FOR UPDATE").First(&lockAddr, addr.ID).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	condition, err := apd.BaseContext.Sub(lockAddr.FreezeDebt, lockAddr.FreezeDebt, t.Amount)
	if err != nil {
		tx.Rollback()
		return err
	}
	if condition.Any() {
		tx.Rollback()
		return errors.New(condition.String())
	}
	err = tx.Save(&lockAddr).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	t.Status = StatusDone
	err = tx.Save(&t).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
