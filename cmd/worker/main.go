package main

import (
	"log"
	"time"

	"github.com/Covertness/ally/pkg/config"
	"github.com/Covertness/ally/pkg/transaction"

	"github.com/Covertness/ally/pkg/ethclient"
	"github.com/Covertness/ally/pkg/etherscan"

	"github.com/Covertness/ally/pkg/hdwallet"
	"github.com/Covertness/ally/pkg/messagequeue"
	"github.com/Covertness/ally/pkg/order"
	"github.com/Covertness/ally/pkg/storage"
	"github.com/beanstalkd/go-beanstalk"
)

func main() {
	db, err := storage.InitDB(config.GetDBDialect(), config.GetDBConnectArgs())
	if err != nil {
		log.Fatalf("db init err: %v", err)
		return
	}
	defer db.Close()

	err = messagequeue.InitQueue()
	if err != nil {
		log.Fatalf("message queue init err: %v", err)
		return
	}

	mnemonic, err := config.GetMnemonic()
	if err != nil {
		log.Fatalf("get mnemonic err: %v", err)
		return
	}

	_, err = hdwallet.Init(mnemonic)
	if err != nil {
		log.Fatalf("hdwallet init err: %v", err)
		return
	}

	_ = etherscan.Init()

	_, err = ethclient.Init("https://" + config.GetInfuraEndpoint() + ".infura.io/v3/" + config.GetInfuraID())
	if err != nil {
		log.Fatalf("ethclient init err: %v", err)
		return
	}

	for {
		id, message, err := messagequeue.GetMessage(5 * time.Second)
		if err != nil {
			cerr, ok := err.(beanstalk.ConnError)
			if ok && cerr.Err == beanstalk.ErrTimeout {
				continue
			}

			log.Println("beanstalk reserve err: ", err)
			return
		}

		log.Println("new message: ", id)

		switch message.Package {
		case "order":
			err = order.JobDispatch(message)
		case "transaction":
			err = transaction.JobDispatch(message)
		default:
			log.Fatalf("unknown package: %s", message.Package)
		}
		if err != nil {
			log.Println("dispatch err: ", err)
			_ = messagequeue.PutbackMessage(id)
			continue
		}

		log.Println("process message success: ", id)

		_ = messagequeue.DeleteMessage(id)

		time.Sleep(1 * time.Second) // avoid api rate limit
	}
}
