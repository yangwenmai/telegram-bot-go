package main

import (
	"log"
	"os"
	"strings"

	"github.com/developer-learning/telegram-bot-go/command"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("night_reading_go_bot"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we should leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		log.Printf("update.Message:%v", update.Message)
		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = `type /sayhi
				/trending go today or /trending go weekly or /trending go monthly
				/status`
		case "sayhi":
			msg.Text = "Hi :)"
		case "status":
			msg.Text = "I'm ok."
		case "trending":
			t := strings.Split(update.Message.Text, " ")
			if strings.TrimSpace(t[1]) == "" {
				msg.Text = "Please input correct language!!!"
			} else {
				since := "today"
				if len(t) >= 3 && t[2] != "" {
					since = t[2]
				}
				msg.Text = command.ListGithubTrending(t[1], since)
				msg.ReplyToMessageID = update.Message.MessageID
				msg.ParseMode = "markdown"
			}
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
