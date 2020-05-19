package main

import (
	"bytes"
	"fmt"
	"github.com/Covertness/ally/pkg/account"
	"github.com/Covertness/ally/pkg/favorite"
	"io"
	"strings"
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

var (
	myFtx *ftx.FTX
	myCoinBase *coinbase.CoinBase
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
		&account.Account{}, &favorite.Favorite{},
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

	myFtx = ftx.Init()
	myCoinBase = coinbase.Init()

	b.Handle("/start", func(m *tb.Message) {
		_, _ = sendResponse(m.Sender, fmt.Sprintf("欢迎 %s %s\n/markets 查看行情", m.Sender.FirstName, m.Sender.LastName))
	})

	b.Handle("/markets", func(m *tb.Message) {
		allMarkets, err := marketPkg.GetAll()
		if err != nil {
			log.Error(err)
			return
		}

		data, err := outputMarkets(allMarkets)
		if err != nil {
			log.Error(err)
			return
		}

		markdownStr := fmt.Sprintf("```\n%s\n```/market 查看详情", data)
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

	b.Handle("/pin", func(m *tb.Message) {
		myAccount, err := account.GetOrCreate(fmt.Sprintf("%d", m.Sender.ID))
		if err != nil {
			log.Error(err)
			return
		}

		typeValue := strings.Split(m.Payload, " ")
		if len(typeValue) < 2 {
			_, _ = sendResponse(m.Sender, fmt.Sprintf("请输入要关注的类别及索引，比如\n/pin market BTC-USD"))
			return
		}

		itemID, err := findFavoriteItemID(typeValue)
		if err != nil {
			log.Error(err)
			return
		}
		if itemID == 0 {
			_, _ = sendResponse(m.Sender, fmt.Sprintf("未找到关注对象，请输入正确的类别及索引，比如\n/pin market BTC-USD"))
			return
		}

		err = favorite.SetMyFavorite(&favorite.Favorite{
			AccountID: myAccount.ID,
			ItemType:  typeValue[0],
			ItemID:    itemID,
		})
		if err != nil && !strings.Contains(err.Error(), "pq: duplicate key value violates unique constraint") {
			log.Error(err)
			return
		}

		_, _ = sendResponse(m.Sender, fmt.Sprintf("关注成功，查看关注列表 /fav"))
	})

	b.Handle("/unpin", func(m *tb.Message) {
		myAccount, err := account.GetOrCreate(fmt.Sprintf("%d", m.Sender.ID))
		if err != nil {
			log.Error(err)
			return
		}

		typeValue := strings.Split(m.Payload, " ")
		if len(typeValue) < 2 {
			_, _ = sendResponse(m.Sender, fmt.Sprintf("请输入要取消关注的类别及索引，比如\n/unpin market BTC-USD"))
			return
		}

		itemID, err := findFavoriteItemID(typeValue)
		if err != nil {
			log.Error(err)
			return
		}
		if itemID == 0 {
			_, _ = sendResponse(m.Sender, fmt.Sprintf("未找到关注对象，请输入正确的类别及索引，比如\n/unpin market BTC-USD"))
			return
		}

		err = favorite.DelMyFavorite(&favorite.Favorite{
			AccountID: myAccount.ID,
			ItemType:  typeValue[0],
			ItemID:    itemID,
		})
		if err != nil {
			log.Error(err)
			return
		}

		_, _ = sendResponse(m.Sender, fmt.Sprintf("取消关注成功，查看关注列表 /fav"))
	})

	b.Handle("/fav", func(m *tb.Message) {
		myAccount, err := account.GetOrCreate(fmt.Sprintf("%d", m.Sender.ID))
		if err != nil {
			log.Error(err)
			return
		}

		favs, err := favorite.GetMyFavorites(myAccount.ID)
		if err != nil {
			log.Error(err)
			return
		}

		var markets []*marketPkg.Market
		for _, fav := range favs {
			switch fav.ItemType {
			case favorite.ItemTypeMarket:
				m, err := marketPkg.GetByID(fav.ItemID)
				if err != nil {
					log.Error(err)
					return
				}
				markets = append(markets, m)
			}
		}

		data, err := outputMarkets(markets)
		if err != nil {
			log.Error(err)
			return
		}

		markdownStr := fmt.Sprintf("```\nMarkets:\n%s\n```/unpin 取消关注", data)
		_, _ = sendResponse(m.Sender, markdownStr, tb.ModeMarkdown)
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

func outputMarkets(allMarkets []*marketPkg.Market) (string, error) {
	ftxMarkets, err := myFtx.GetMarkets()
	if err != nil {
		return "", err
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
	return buf.String(), nil
}

func findFavoriteItemID(typeValue []string) (uint, error) {
	switch typeValue[0] {
	case favorite.ItemTypeMarket:
		market, err := marketPkg.GetByName(typeValue[1])
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			return 0, err
		}
		if market == nil || !market.Enable {
			return 0, nil
		}

		return market.ID, nil
	default:
		return 0, nil
	}
}
