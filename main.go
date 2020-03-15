package main

import (
	"fmt"
	"log"
	"os"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var (
		port      = os.Getenv("PORT")
		publicURL = "tg.me/ElonVMaskeBot"
		token     = "1105551961:AAE_gaIQhxz08hiRiI1T8eRPzWXsRfl-DqE"
	)

	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}

	pref := tb.Settings{
		Token:  token,
		Poller: webhook,
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	b.Handle("/hello", func(m *tb.Message) {
		fmt.Println(m.Sender, "joined")
		b.Send(m.Sender, "Hi!")
	})

	b.Start()
}
