package config

import (
	"strings"

	"github.com/tyler-smith/go-bip39"
)

// GetMnemonic get the hdwallet mnemonic
func GetMnemonic() (string, error) {
	mnemonic, err := getconfig("mnemonic")
	if err != nil {
		if !strings.Contains(err.Error(), "record not found") {
			return "", err
		}

		entropy, err := bip39.NewEntropy(256)
		if err != nil {
			return "", err
		}
		mnemonic, err := bip39.NewMnemonic(entropy)
		if err != nil {
			return "", err
		}

		v, _ := encodeValue(map[string]string{
			"mnemonic": mnemonic,
		})

		err = setconfig("mnemonic", v)
		if err != nil {
			return "", err
		}
		return mnemonic, nil
	}

	v, err := decodeValue(mnemonic)
	if err != nil {
		return "", err
	}
	return v["mnemonic"], nil
}
