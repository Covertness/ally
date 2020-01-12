package main

import (
	"log"

	"github.com/Covertness/ally/pkg/account"
	"github.com/Covertness/ally/pkg/address"
	"github.com/Covertness/ally/pkg/config"
	"github.com/Covertness/ally/pkg/etherscan"
	"github.com/Covertness/ally/pkg/transaction"

	"github.com/Covertness/ally/pkg/transactiongroup"

	"github.com/Covertness/ally/pkg/admin"

	"github.com/Covertness/ally/pkg/hdwallet"
	"github.com/Covertness/ally/pkg/item"
	"github.com/Covertness/ally/pkg/order"
	"github.com/Covertness/ally/pkg/stat"
	"github.com/Covertness/ally/pkg/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := storage.InitDB(config.GetDBDialect(), config.GetDBConnectArgs())
	if err != nil {
		log.Fatalf("db init err: %v", err)
		return
	}
	defer db.Close()

	err = db.AutoMigrate(
		&config.Config{},
		&account.Account{}, &item.Item{},
		&address.Address{}, &order.Order{},
		&transaction.Transaction{}, &transactiongroup.TransactionGroup{},
	).Error
	if err != nil {
		log.Fatalf("db auto migrate err: %v", err)
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

	r := gin.Default()
	v1 := r.Group("/api/v1")
	item.Register(v1.Group("/items"))
	order.Register(v1.Group("/orders"))
	stat.Register(v1.Group("/stat"))
	admin.Register(v1.Group("/admin"))
	transactiongroup.Register(v1.Group("/transactiongroups"))

	r.Run() // listen and serve on 0.0.0.0:8080
}
