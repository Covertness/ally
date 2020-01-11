package hdwallet

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

var wallet *hdwallet.Wallet

// Init create a wallet from seed
func Init(mnemonic string) (*hdwallet.Wallet, error) {
	w, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}

	wallet = w
	return wallet, nil
}

// RootAddress get root address
func RootAddress() (*accounts.Account, error) {
	return DeriveAddress("m/0")
}

// AddressByID get address by id
func AddressByID(id uint) (string, *accounts.Account, error) {
	p := fmt.Sprintf("m/%d", id)
	a, err := DeriveAddress(p)
	return p, a, err
}

// AddressPrivateKey get private key by address
func AddressPrivateKey(a *accounts.Account) (*ecdsa.PrivateKey, error) {
	return wallet.PrivateKey(*a)
}

// DeriveAddress derive an address from path
func DeriveAddress(path string) (*accounts.Account, error) {
	p, err := hdwallet.ParseDerivationPath(path)
	if err != nil {
		return nil, err
	}

	a, err := wallet.Derive(p, false)
	return &a, err
}
