package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gsora/nibberbot/nibber"
)

const (
	suh           = "suh my ni\xF0\x9F\x85\xB1\xF0\x9F\x85\xB1a"
	breathingSuh  = "suh my \xF0\x9F\x85\xB1reathing ni\xF0\x9F\x85\xB1\xF0\x9F\x85\xB1a"
	clappingNibba = "clapping ni\xF0\x9F\x85\xB1\xF0\x9F\x85\xB1a"
)

var (
	nibberInstance nibber.Nibber
	apiKey         string
	certPath       string
	keyPath        string
	debug          bool
	domain         string
	port           string
)

func main() {
	setupParams()

	nibberInstance = nibber.NewNibber(nibber.Emojis)

	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.RemoveWebhook()
	if err != nil {
		log.Fatal(err)
	}
	_, err = bot.SetWebhook(tgbotapi.NewWebhook("https://" + domain + ":" + port + "/" + apiKey))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/" + apiKey)
	go startServer()
	for update := range updates {
		// handle every update in a separate goroutine
		go handleUpdate(update, bot)
	}

}

func setupParams() {
	flag.StringVar(&certPath, "cert", "", "required, TLS certificate path")
	flag.StringVar(&keyPath, "key", "", "required, TLS key path")
	flag.StringVar(&apiKey, "apikey", "", "required, Telegram bot API key")
	flag.StringVar(&domain, "domain", "", "required, domain associated to the TLS cert+key and the server where this bot will be running")
	flag.StringVar(&port, "port", "88", "port to run on, must be 443, 80, 88, 8443")
	flag.BoolVar(&debug, "debug", false, "debug Telegram bot interactions")
	flag.Parse()

	if certPath == "" || keyPath == "" || apiKey == "" || domain == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func startServer() {
	go log.Fatal(http.ListenAndServeTLS("0.0.0.0:"+port, certPath, keyPath, nil))
}
