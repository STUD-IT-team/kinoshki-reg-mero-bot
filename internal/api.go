package api

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Button struct {
	Text   string
	Data   string
	URL    string
	WebApp *tgbotapi.WebAppInfo
}

type ReplyButton struct {
	Text      string
	IsContact bool
	WebApp    *tgbotapi.WebAppInfo
}

func SendButtons(bot *tgbotapi.BotAPI, chatID int64, text string, buttons [][]Button) error {
	var rows [][]tgbotapi.InlineKeyboardButton

	for _, btnRow := range buttons {
		var row []tgbotapi.InlineKeyboardButton
		for _, btn := range btnRow {
			var kb tgbotapi.InlineKeyboardButton
			switch {
			case btn.URL != "":
				kb = tgbotapi.NewInlineKeyboardButtonURL(btn.Text, btn.URL)
			case btn.WebApp != nil:
				kb = tgbotapi.NewInlineKeyboardButtonWebApp(btn.Text, *btn.WebApp)
			default:
				kb = tgbotapi.NewInlineKeyboardButtonData(btn.Text, btn.Data)
			}
			row = append(row, kb)
		}
		rows = append(rows, row)
	}

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(rows...)
	_, err := bot.Send(msg)
	return err
}

func SendReplyKeyboard(bot *tgbotapi.BotAPI, chatID int64, text string, buttons [][]ReplyButton, oneTime bool) error {
	var keyboard [][]tgbotapi.KeyboardButton

	for _, row := range buttons {
		var kbRow []tgbotapi.KeyboardButton
		for _, btn := range row {
			var kb tgbotapi.KeyboardButton
			switch {
			case btn.IsContact:
				kb = tgbotapi.NewKeyboardButtonContact(btn.Text)
			case btn.WebApp != nil:
				kb = tgbotapi.NewKeyboardButtonWebApp(btn.Text, *btn.WebApp)
			default:
				kb = tgbotapi.NewKeyboardButton(btn.Text)
			}
			kbRow = append(kbRow, kb)
		}
		keyboard = append(keyboard, kbRow)
	}

	msg := tgbotapi.NewMessage(chatID, text)
	replyMarkup := tgbotapi.NewReplyKeyboard(keyboard...)
	replyMarkup.OneTimeKeyboard = oneTime
	msg.ReplyMarkup = replyMarkup
	_, err := bot.Send(msg)
	return err
}
