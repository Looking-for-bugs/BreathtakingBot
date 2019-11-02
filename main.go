package main

import (
	"net/http"
	"log"
	"os"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	token := os.Getenv("BOTTOKEN")
	if token == "" {
		log.Fatal("bot token is not provided")
	}

	serverURL := os.Getenv("URL")
	port := os.Getenv("PORT")
	if serverURL == "" {
		log.Fatal("server url is not provided")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal("error while initializing bot ", err)
	}

	log.Printf("Authorized as %s", bot.Self.FirstName)


	hookURL := fmt.Sprintf("%s:%s/%s", serverURL, port, token)
	tgbotapi.NewWebhook(hookURL)
	
	updates := bot.ListenForWebhook("/" + token)
	go http.ListenAndServe(":" + port, nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}
}
