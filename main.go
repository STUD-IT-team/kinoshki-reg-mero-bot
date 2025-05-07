package main

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var userStates = make(map[int64]string)

func main() {
	bot, err := tgbotapi.NewBotAPI("7767177359:AAEGm6K_QQXHHVYc9ashtu__qO9llCjWpqI")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	commands := []tgbotapi.BotCommand{
		{Command: "start", Description: "Начало работы"},
		{Command: "register", Description: "Регистрация"},
		{Command: "schedule", Description: "Расписание"},
		{Command: "elements", Description: "Графические элементы"},
	}
	_, err = bot.Request(tgbotapi.NewSetMyCommands(commands...))
	if err != nil {
		log.Fatal(err)
	}

	u := tgbotapi.NewUpdate(0)
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery != nil {
			handleCallback(bot, update)
			continue
		}

		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID
		msgText := update.Message.Text

		switch {
		case msgText == "/start":
			msg := tgbotapi.NewMessage(chatID, "Добро пожаловать! Используйте команды:\n/elements - демо графических элементов\n/register - регистрация\n/schedule - расписание")
			bot.Send(msg)

		case msgText == "/register":
			userStates[chatID] = "awaiting_name"
			msg := tgbotapi.NewMessage(chatID, "Введите ваше ФИО:")
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			bot.Send(msg)

		case msgText == "/schedule":
			msg := tgbotapi.NewMessage(chatID, "Расписание:")
			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonURL("Открыть расписание", "https://bmstu.ru/schedule"),
				),
			)
			bot.Send(msg)

		case msgText == "/elements":
			showElementsMenu(bot, chatID)

		default:
			if state, ok := userStates[chatID]; ok {
				switch state {
				case "awaiting_name":
					userStates[chatID] = "awaiting_group"
					msg := tgbotapi.NewMessage(chatID, "Теперь выберите факультет:")
					msg.ReplyMarkup = facultyKeyboard()
					bot.Send(msg)

				case "awaiting_group":
					msg := tgbotapi.NewMessage(chatID, "Пожалуйста, выберите факультет из кнопок ниже:")
					msg.ReplyMarkup = facultyKeyboard()
					bot.Send(msg)
				}
			}
		}
	}
}

func showElementsMenu(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Выберите графический элемент для демонстрации:")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Reply-клавиатура", "show_reply"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Inline-кнопки", "show_inline"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Кнопка контакта", "show_contact"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Кнопка геолокации", "show_location"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("WebApp кнопка", "show_webapp"),
		),
	)
	bot.Send(msg)
}

func handleCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	query := update.CallbackQuery
	chatID := query.Message.Chat.ID
	data := query.Data

	switch {
	case data == "btn1":
		msg := tgbotapi.NewMessage(chatID, "Вы нажали кнопку 1")
		bot.Send(msg)

	case data == "btn2":
		msg := tgbotapi.NewMessage(chatID, "Вы нажали кнопку 2")
		bot.Send(msg)

	case data == "show_reply":
		msg := tgbotapi.NewMessage(chatID, "Пример Reply-клавиатуры:")
		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Кнопка 1"),
				tgbotapi.NewKeyboardButton("Кнопка 2"),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Кнопка 3"),
			),
		)
		bot.Send(msg)

	case data == "show_inline":
		msg := tgbotapi.NewMessage(chatID, "Пример Inline-кнопок:")
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Кнопка 1", "btn1"),
				tgbotapi.NewInlineKeyboardButtonData("Кнопка 2", "btn2"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonURL("Открыть сайт", "https://bmstu.ru/"),
			),
		)
		bot.Send(msg)

	case data == "show_contact":
		msg := tgbotapi.NewMessage(chatID, "Пример кнопки для отправки контакта (клавиатура автоматически закроется после отправки):")
		keyboard := tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButtonContact("Отправить мой контакт"),
			),
		)
		keyboard.OneTimeKeyboard = true
		msg.ReplyMarkup = keyboard
		bot.Send(msg)

	case data == "show_location":
		msg := tgbotapi.NewMessage(chatID, "Пример кнопки для отправки геолокации (клавиатура автоматически закроется после отправки):")
		keyboard := tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButtonLocation("Отправить мою локацию"),
			),
		)
		keyboard.OneTimeKeyboard = true
		msg.ReplyMarkup = keyboard
		bot.Send(msg)

	case data == "show_webapp":
		msg := tgbotapi.NewMessage(chatID, "Пример WebApp кнопки:")
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonWebApp(
					"Открыть WebApp",
					tgbotapi.WebAppInfo{URL: "https://telegram.org"},
				),
			),
		)
		bot.Send(msg)

	case strings.HasPrefix(data, "faculty_"):
		keyboard := groupKeyboard()
		editMsg := tgbotapi.NewEditMessageText(
			chatID,
			query.Message.MessageID,
			"Выберите группу:",
		)
		editMsg.ReplyMarkup = &keyboard
		bot.Send(editMsg)

	case strings.HasPrefix(data, "group_"):
		delete(userStates, chatID)
		editMsg := tgbotapi.NewEditMessageText(
			chatID,
			query.Message.MessageID,
			"Регистрация завершена! ✅",
		)
		bot.Send(editMsg)
	}
}

func facultyKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("ИУ5", "faculty_iu5"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("ИУ6", "faculty_iu6"),
		),
	)
}

func groupKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("ИУ5-11Б", "group_iu5_11b"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("ИУ6-12Б", "group_iu6_12b"),
		),
	)
}
