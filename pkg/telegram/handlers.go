package telegram

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	gowiki "github.com/trietmn/go-wiki"
)

const commandStart = "start"
const new_lang = "ru"

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Вопросы | Помощь в чем-то"),
		tgbotapi.NewKeyboardButton("Отправить отзыв | вопрос автору"),
		tgbotapi.NewKeyboardButton("Мой автор"),
	),
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Такой команды нету :(")

	switch message.Command() {
	case commandStart:
		msg.Text = `Привет, это бот "помощник | психолог и т.д" помогает найти информацию и помочь в чем то. В конце своего вопроса всегда добавляйте знак вопрос "?".
		`
		msg.ReplyMarkup = numericKeyboard
		_, err := b.bot.Send(msg)
		return err
	default:
		_, err := b.bot.Send(msg)
		return err
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	// log.Printf("[%s] %s", message.From.UserName, message.Text)
	if message.Text == "Вопросы | Помощь в чем-то" {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Напишите свой вопрос!")
		b.bot.Send(msg)
	} else if strings.ContainsAny(message.Text, "?") {
		from_msg := strings.ReplaceAll(message.Text, "?", "")
		gowiki.SetLanguage(new_lang)
		res, err := gowiki.Summary(from_msg, 5, -1, false, true)
		if err != nil {
			log.Panic(err)
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, res)
		b.bot.Send(msg)
	} else if message.Text != "Вопросы | Помощь в чем-то" {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Loading...")
		b.bot.Send(msg)
	}
	switch message.Text {
	case "Отправить отзыв | вопрос автору":
		msg := tgbotapi.NewMessage(message.Chat.ID, "Напишите отзыв или вопрос, а я отправлю его автору.")
		b.bot.Send(msg)
	case "Мой автор":
		msg := tgbotapi.NewMessage(message.Chat.ID, "Автор этого бота: @oxxqkk")
		b.bot.Send(msg)
	default:
		fmt.Println("worked")
	}
}
