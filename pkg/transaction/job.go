package transaction

import (
	"log"
	"math/big"

	"github.com/Covertness/ally/pkg/etherscan"
	"github.com/Covertness/ally/pkg/messagequeue"
	"github.com/Covertness/ally/pkg/storage"

	"github.com/bamzi/jobrunner"
)

// JobSchedule called by dispatcher
func JobSchedule() error {
	err := jobrunner.Schedule("@every 10s", processTransactionsSchedule{})
	if err != nil {
		return err
	}
	return nil
}

// JobDispatch called by worker
func JobDispatch(msg *messagequeue.Message) error {
	params := msg.Params

	switch msg.Action {
	case "processTransaction":
		return processTransaction(uint(params["transactionID"].(float64)))
	default:
		log.Fatalf("unknown action %s", msg.Action)
		return nil
	}
}

type processTransactionsSchedule struct{}

func (p processTransactionsSchedule) Run() {
	transactions, err := GetByStatuses([]string{
		StatusInit, StatusPending,
	})
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	for _, t := range transactions {
		_, err = messagequeue.SendMessage(&messagequeue.Message{
			Package: "transaction",
			Action:  "processTransaction",
			Params: map[string]interface{}{
				"transactionID": t.ID,
			},
		})
		if err != nil {
			log.Fatalf("%v", err)
		}
	}
}

func processTransaction(transactionID uint) error {
	var t Transaction
	err := storage.GetDB().Preload("Address").First(&t, transactionID).Error
	if err != nil {
		return err
	}

	switch t.Status {
	case StatusInit:
		return provision(&t)
	case StatusPending:
		return checkSuccess(&t)
	default:
		log.Fatalf("unexpected status %s", t.Status)
		return nil
	}
}

func provision(t *Transaction) error {
	return t.RequestApply()
}

func checkSuccess(t *Transaction) error {
	tx, err := etherscan.GetInstance().GetTransaction(t.Hash)
	if err != nil {
		return err
	}
	blockNumber, err := etherscan.GetInstance().GetBlockNumber()
	if err != nil {
		return err
	}

	wantedConfirmations := big.NewInt(15)
	var confirmations big.Int
	if tx.BlockNumber == nil {
		return nil
	}
	if confirmations.Sub(blockNumber, tx.BlockNumber).Cmp(wantedConfirmations) < 0 {
		return nil
	}

	return t.Apply()
}
