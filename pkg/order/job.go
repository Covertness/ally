package order

import (
	"errors"
	"log"
	"time"

	"github.com/Covertness/ally/pkg/etherscan"
	"github.com/Covertness/ally/pkg/messagequeue"
	"github.com/Covertness/ally/pkg/storage"

	"github.com/bamzi/jobrunner"
	"github.com/cockroachdb/apd"
)

// JobSchedule called by dispatcher
func JobSchedule() error {
	err := jobrunner.Schedule("@every 10s", processOrdersSchedule{})
	if err != nil {
		return err
	}
	return nil
}

// JobDispatch called by worker
func JobDispatch(msg *messagequeue.Message) error {
	params := msg.Params

	switch msg.Action {
	case "processOrder":
		return processOrder(uint(params["orderID"].(float64)))
	default:
		log.Fatalf("unknown action %s", msg.Action)
		return nil
	}
}

type processOrdersSchedule struct{}

func (p processOrdersSchedule) Run() {
	orders, err := GetByStatuses([]string{
		StatusInit, StatusWaitingAddress, StatusDepositable, StatusDepositing,
	})
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	for _, o := range orders {
		_, err = messagequeue.SendMessage(&messagequeue.Message{
			Package: "order",
			Action:  "processOrder",
			Params: map[string]interface{}{
				"orderID": o.ID,
			},
		})
		if err != nil {
			log.Fatalf("%v", err)
		}
	}
}

func processOrder(orderID uint) error {
	var o Order
	err := storage.GetDB().Preload("Account").Preload("Address").Preload("Item").First(&o, orderID).Error
	if err != nil {
		return err
	}

	switch o.Status {
	case StatusInit, StatusWaitingAddress:
		return provision(&o)
	case StatusDepositable, StatusDepositing:
		return checkDeposit(&o)
	case StatusDone, StatusTimeout:
		return nil
	default:
		log.Fatalf("unexpected status %s", o.Status)
		return nil
	}
}

func provision(o *Order) error {
	return o.RequestAddress()
}

func checkDeposit(o *Order) error {
	if o.UpdatedAt.Add(time.Hour).Before(time.Now()) {
		err := storage.GetDB().Model(o).Updates(Order{
			Status: StatusTimeout,
		}).Error
		if err != nil {
			return err
		}
	}

	balance, err := etherscan.GetInstance().GetTokenBalance(o.Address.Address)
	if err != nil {
		return err
	}

	if balance.IsZero() {
		return nil
	}

	if o.Status == StatusDepositable {
		err = storage.GetDB().Model(o).Updates(Order{
			Status: StatusDepositing,
		}).Error
		if err != nil {
			return err
		}
	}

	scaner := etherscan.GetInstance()
	decimals, err := scaner.GetTokenDecimals()
	if err != nil {
		return err
	}

	condition, err := apd.BaseContext.Mul(balance, balance, apd.New(1, -int32(decimals)))
	if err != nil {
		return err
	}
	if condition.Any() {
		return errors.New(condition.String())
	}

	price, err := o.GetItemPrice()
	if err != nil {
		return err
	}

	if balance.Cmp(price) < 0 {
		return nil
	}

	return o.ApplyDeposit(balance)
}
