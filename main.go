package main

import (
	"log"
	"github.com/Syfaro/telegram-bot-api"
)


func main() {
	bot, err := tgbotapi.NewBotAPI("1654399226:AAHfx2-TX61Zvxu5lOtJiwYn2TMH8rVzJpI")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
}
