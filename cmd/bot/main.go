package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/luism6n/jumblebot/jumble"

	"gopkg.in/telegram-bot-api.v4"
)

var (
	token *string
	debug *bool
)

func main() {
	bot, updates := setupBot()

	for update := range updates {
		if update.InlineQuery == nil {
			continue
		}

		jumbledText := jumble.Jumble(update.InlineQuery.Query)

		result := newInlineQueryResultArticle(jumbledText)
		config := newInlineConfig(update.InlineQuery.ID, result)

		res, err := bot.AnswerInlineQuery(config)
		if err != nil {
			log.Printf("Error:\nerr: %s\nres: %+v\nquery: %s\njumbled text: %s", err.Error(), res, update.InlineQuery.Query, jumbledText)
		}
	}
}

func setupBot() (*tgbotapi.BotAPI, tgbotapi.UpdatesChannel) {
	readCommandLineFlags()

	bot, err := tgbotapi.NewBotAPI(*token)
	if err != nil {
		fmt.Printf("Bot token is invalid: %s\n", *token)
		os.Exit(1)
	}

	bot.Debug = *debug

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Never returns error
	updates, _ := bot.GetUpdatesChan(u)

	return bot, updates
}

func readCommandLineFlags() {
	debug = flag.Bool("debug", false, "If the bot should run in debug mode")
	token = flag.String("token", "No token provided", "The bot token")
	flag.Parse()
}

func newInlineQueryResultArticle(text string) tgbotapi.InlineQueryResultArticle {
	return tgbotapi.InlineQueryResultArticle{
		Type:        "article",
		ID:          "only result",
		Title:       "Jumbled text",
		Description: text,
		InputMessageContent: tgbotapi.InputTextMessageContent{
			Text: text,
		},
	}
}

func newInlineConfig(queryID string, onlyResult tgbotapi.InlineQueryResultArticle) tgbotapi.InlineConfig {
	return tgbotapi.InlineConfig{
		InlineQueryID: queryID,
		Results:       castToInterfaceSlice([]tgbotapi.InlineQueryResultArticle{onlyResult}),
		CacheTime:     300,
	}
}

func castToInterfaceSlice(iqra []tgbotapi.InlineQueryResultArticle) []interface{} {
	s := make([]interface{}, len(iqra))
	for i, v := range iqra {
		s[i] = v
	}

	return s
}
