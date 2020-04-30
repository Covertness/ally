package main

import (
	"fmt"
	"github.com/Covertness/ally/pkg/ftx"
	"log"
	"time"

	"github.com/Covertness/ally/pkg/config"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  config.GetTelegramToken(),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	myFtx := ftx.Init()

	b.Handle("/market", func(m *tb.Message) {
		market, err := myFtx.GetMarket(m.Payload)
		if err != nil {
			log.Printf("error %v", err)
			return
		}

		_, _ = b.Send(m.Sender, fmt.Sprintf("last %f", market.Last))
	})

	log.Println("bot is working...")
	b.Start()
}
