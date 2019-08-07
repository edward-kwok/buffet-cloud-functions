package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	pubsub "cloud.google.com/go/pubsub"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Webhook(w http.ResponseWriter, r *http.Request) {
	const topic = "buffet"
	const proj = "buffet-edward-247815"
	token := os.Getenv("TOKEN")
	bytes, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	var update tgbotapi.Update
	json.Unmarshal(bytes, &update)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()

	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	fmt.Printf("%+v\n", update.Message)
	if update.Message == nil { // ignore any non-Message updates
		return
	}

	if !update.Message.IsCommand() { // ignore any non-command Messages
		return
	}
	switch update.Message.Command() {
	case "buffet":
		fmt.Println("Command buffet is triggered")
		ctx := context.Background()
		client, err := pubsub.NewClient(ctx, proj)
		if err != nil {
			log.Fatal(err)
		}
		publish(client, topic)
	}

}

func publish(client *pubsub.Client, topic string) error {
	ctx := context.Background()
	t := client.Topic(topic)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte("Hello world!"),
		Attributes: map[string]string{
			"origin":   "golang",
			"username": "gcp",
		},
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
	return nil
}
