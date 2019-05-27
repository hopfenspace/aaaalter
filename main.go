package main

import (
	"log"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"regexp"
	"bytes"
	"strings"
	"io/ioutil"
)

func main() {
	token, _ := ioutil.ReadFile("token.txt")
	bot, err := tgbotapi.NewBotAPI(string(token))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	r, _ := regexp.Compile("(?i)[a]+lter")

	for update := range updates {
		if update.Message == nil {
			continue
		}

		match := r.MatchString(update.Message.Text)
		
		if match {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			var buffer bytes.Buffer

			alter := r.FindAllString(update.Message.Text, 1)
			acount := len(alter[0]) - 4

			buffer.WriteString(strings.Repeat("a", acount * 2))
			buffer.WriteString("lter")

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, buffer.String())
			msg.ReplyToMessageID = update.Message.MessageID
	
			bot.Send(msg)	
		}
	}
}