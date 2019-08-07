package cmd

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// PubSubMessage is the payload of a Pub/Sub event.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

func Trigger(ctx context.Context, m PubSubMessage) error {
	token := os.Getenv("TOKEN")
	resp, err := http.Get("https://www.buys.hk/best-offers-2019/grand-hyatt/html/tiffin-dinner-checkout-zh.html")
	if err != nil {
		return err
	}
	re := regexp.MustCompile(`4:.\[\s*\]`)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	bs := string(body)
	dates := re.FindString(bs)

	if dates == "" {
		bot, err := tgbotapi.NewBotAPI(token)
		if err != nil {
			log.Panic(err)
		}
		bot.Debug = true
		log.Printf("Authorized on account %s", bot.Self.UserName)
		log.Printf("Buffet Ready")
		m := tgbotapi.NewMessage(-1001279698998, "Hello, Buffet ready")
		bot.Send(m)
	} else {
		log.Printf("No buffet at " + time.Now().Format(time.RFC850))
	}
	return nil

}
