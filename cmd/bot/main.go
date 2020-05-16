package main

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/Covertness/ally/pkg/coinbase"
	"github.com/Covertness/ally/pkg/config"
	"github.com/Covertness/ally/pkg/ftx"
	marketPkg "github.com/Covertness/ally/pkg/market"
	"github.com/Covertness/ally/pkg/storage"
	"github.com/jinzhu/gorm"
	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetLevel(log.DebugLevel)

	db, err := storage.InitDB(config.GetDBDialect(), config.GetDBConnectArgs())
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() {
		_ = db.Close()
	}()

	err = db.AutoMigrate(
		&marketPkg.Market{},
	).Error
	if err != nil {
		log.Fatal(err)
		return
	}

	logPoller := tb.NewMiddlewarePoller(&tb.LongPoller{Timeout: 10 * time.Second}, func(upd *tb.Update) bool {
		if upd.Message != nil {
			log.WithFields(log.Fields{"type": "request", "from": upd.Message.Sender.Username}).Debug(upd.Message.Text)
		}
		return true
	})
	b, err := tb.NewBot(tb.Settings{
		Token:  config.GetTelegramToken(),
		Poller: logPoller,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	sendResponse := func(to *tb.User, what interface{}, options ...interface{}) (*tb.Message, error) {
		log.WithFields(log.Fields{"type": "response", "to": to.Username}).Debug(what)
		return b.Send(to, what, options...)
	}

	myFtx := ftx.Init()
	myCoinBase := coinbase.Init()

	b.Handle("/start", func(m *tb.Message) {
		_, _ = sendResponse(m.Sender, fmt.Sprintf("欢迎 %s %s\n/markets 查看行情", m.Sender.FirstName, m.Sender.LastName))
	})

	b.Handle("/markets", func(m *tb.Message) {
		allMarkets, err := marketPkg.GetAll()
		if err != nil {
			log.Error(err)
			return
		}

		ftxMarkets, err := myFtx.GetMarkets()
		if err != nil {
			log.Error(err)
			return
		}

		var buf bytes.Buffer
		table := newTable(&buf)

		for _, market := range allMarkets {
			switch market.Provider {
			case marketPkg.ProviderFTX:
				m := funk.Find(ftxMarkets, func(m *ftx.Market) bool { return m.Name == market.Name })
				if m != nil {
					ftxMarket := m.(*ftx.Market)
					table.Append([]string{market.Name, fmt.Sprintf("%f", ftxMarket.Last)})
				}
			case marketPkg.ProviderCoinBase:
				p, err := myCoinBase.GetPriceSpot(market.Name)
				if err != nil {
					log.Error(err)
					continue
				}
				table.Append([]string{market.Name, p.Amount})
			}
		}
		table.Render()

		markdownStr := fmt.Sprintf("```\n%s\n```/market 查看详情", buf.String())
		_, _ = sendResponse(m.Sender, markdownStr, tb.ModeMarkdown)
	})

	b.Handle("/market", func(m *tb.Message) {
		marketName := m.Payload
		if len(marketName) == 0 {
			_, _ = sendResponse(m.Sender, fmt.Sprintf("请输入要查询行情的名称，比如\n/market BTC-USD"))
			return
		}

		market, err := marketPkg.GetByName(marketName)
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			log.Error(err)
			return
		}
		if market == nil || !market.Enable {
			_, _ = sendResponse(m.Sender, fmt.Sprintf("未找到行情，请输入正确的名称，比如\n/market BTC-USD"))
			return
		}

		var price string
		switch market.Provider {
		case marketPkg.ProviderFTX:
			ftxMarket, err := myFtx.GetMarket(market.Name)
			if err != nil {
				log.Error(err)
				return
			}
			price = fmt.Sprintf("%f", ftxMarket.Last)
		case marketPkg.ProviderCoinBase:
			p, err := myCoinBase.GetPriceSpot(market.Name)
			if err != nil {
				log.Error(err)
				return
			}
			price = p.Amount
		}

		_, _ = sendResponse(m.Sender, fmt.Sprintf("名称：\n%s", market.Name))
		_, _ = sendResponse(m.Sender, fmt.Sprintf("简介：\n%s", market.Description))
		_, _ = sendResponse(m.Sender, fmt.Sprintf("最近成交价：\n%s", price))
	})

	log.Info("bot is working...")
	b.Start()
}

func newTable(writer io.Writer) *tablewriter.Table {
	table := tablewriter.NewWriter(writer)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	return table
}
