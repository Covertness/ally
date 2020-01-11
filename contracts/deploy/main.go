package main

import (
	"log"

	"github.com/Covertness/ally/contracts"
	"github.com/Covertness/ally/pkg/config"
	"github.com/Covertness/ally/pkg/ethclient"
	"github.com/Covertness/ally/pkg/hdwallet"
	"github.com/Covertness/ally/pkg/storage"
)

func main() {
	db, err := storage.InitDB(config.GetDBDialect(), config.GetDBConnectArgs())
	if err != nil {
		log.Fatalf("db init err: %v", err)
		return
	}
	defer db.Close()

	client, err := ethclient.Init("https://" + config.GetInfuraEndpoint() + ".infura.io/v3/" + config.GetInfuraID())
	if err != nil {
		log.Fatal(err)
	}

	mnemonic, err := config.GetMnemonic()
	if err != nil {
		log.Fatalf("get mnemonic err: %v", err)
		return
	}

	_, err = hdwallet.Init(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	root, err := hdwallet.RootAddress()
	if err != nil {
		log.Fatal(err)
	}

	nonce, err := client.GenNonce()
	if err != nil {
		log.Fatal(err)
	}

	auth, err := client.GetSigner(nonce)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("from address: ", root.Address.Hex())

	address, tx, instance, err := contracts.DeployAddressFactory(auth, client)
	if err != nil {
		log.Println("please make sure the from address has enough funds")
		log.Fatal(err)
	}

	log.Println("AddressFactory contract address: ", address.Hex())
	log.Println("AddressFactory deploy transaction: ", tx.Hash().Hex())

	_ = instance
}
