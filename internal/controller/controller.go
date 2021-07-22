package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"tezos_node_exporter/internal/header"
)

type controller struct {
	id     int64
	local  string
	remote string
	bot    *tgbotapi.BotAPI
}

type Controller interface {
	Run(ctx context.Context)
}

func getToken() string {
	if len(os.Args) > 1 {
		log.Println("got token from command line arg")
		return os.Args[1]
	}
	v := os.Getenv("BOT_TOKEN")
	if v != "" {
		log.Println("got token from envvar")
		return v
	}
	log.Fatal("token not set. set it as commandline arg or in BOT_TOKEN envvar")
	return ""
}

func getheader(url string) (h *header.Header, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	h = new(header.Header)
	err = json.NewDecoder(resp.Body).Decode(h)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return h, nil
}

func (c *controller) Run(ctx context.Context) {
	appError := tgbotapi.NewMessage(c.id, "App error")
	for {
		h1, err := getheader(c.local)
		if err != nil {
			c.bot.Send(appError)
			log.Printf("Error acessing the %s", c.local)
			return
		}
		h2, err := getheader(c.remote)
		if err != nil {
			c.bot.Send(appError)
			log.Printf("Error acessing the %s", c.remote)
			return
		}
		if h1.Timestamp.Sub(h2.Timestamp).Milliseconds() > time.Minute.Milliseconds() {
			// alert takes place
			m := tgbotapi.NewMessage(c.id, "desync detected")
			_, err = c.bot.Send(m)
			if err != nil {
				log.Println("Error sending the message")
			}
			log.Println("Some desync detected")
		}
	}
}

func NewController(local, remote string, id int64) Controller {
	bot, _ := tgbotapi.NewBotAPI(getToken())
	return &controller{local: local, remote: remote, bot: bot, id: id}
}
