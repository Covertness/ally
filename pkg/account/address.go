package account

import (
	"log"
	"math/big"

	"github.com/jinzhu/gorm"

	"github.com/Covertness/ally/pkg/address"
	"github.com/Covertness/ally/pkg/ethclient"
)

// RequestAddress create an address belong to the account
func (a *Account) RequestAddress(tx *gorm.DB) (*address.Address, error) {
	newAddress := &address.Address{
		AccountID: a.ID,
	}
	err := tx.Model(a).Association("Addresses").Append(newAddress).Error
	if err != nil {
		return nil, err
	}

	eth := ethclient.GetInstance()
	factory, err := eth.GetAddressFactory()
	if err != nil {
		return nil, err
	}

	nonce, err := eth.GenNonce()
	if err != nil {
		return nil, err
	}

	signer, err := eth.GetSigner(nonce)
	if err != nil {
		return nil, err
	}

	id := new(big.Int).SetUint64(uint64(newAddress.ID))

	t, err := factory.CreateReceivers(signer, id)
	if err != nil {
		return nil, err
	}

	log.Println(t.Hash().String())

	return newAddress, nil
}
