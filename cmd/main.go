package main

import (
	"fmt"
	"log"

	"github.com/Staspnm/telegrambot/internal/adapters/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	db := postgres.New()

	usrs, err := db.GetUser()

	for _, usr := range usrs {
		fmt.Printf("%d, %s, %s, %d, %s, %s, %s, %s, %s, %s, %s",
			usr.User_pk,
			usr.User_uuid,
			usr.Username,
			usr.Role_id,
			usr.Usertg,
			usr.Enabled,
			usr.Phone,
			usr.Dop_info,
			usr.Last_connection_time,
			usr.Last_network_address,
			usr.Update_time)
	}
	// конец запроса

	// добавление данных в BD

	// _, err = db.Exec("INSERT INTO rote.users (user_uuid, username, role_id, usertg, phone, dop_info, last_connection_time, last_network_address, update_time) VALUES (?, ?, ?, ?, ?, ?, now() at time zone 'utc', ?, now() at time zone 'utc')", 2, "MolEliza", 1, "popan", "89124556998", "45", "hjgb")
	// if err != nil {
	// 	fmt.Println("Ошибка добавления данных", err)
	// 	log.Fatal(err)
	// }
	// 	_, err = db.Exec(`
	//     INSERT INTO rote.users (
	//         user_uuid, username, role_id, usertg, phone, dop_info,
	//         last_connection_time, last_network_address, update_time
	//     ) VALUES ($1, $2, $3, $4, $5, $6, now() at time zone 'utc', $7, now() at time zone 'utc')
	// `, 2, "MolEliza", 1, "popan", "89124556998", "45", "hjgb")
	// 	if err != nil {
	// 		fmt.Println("Ошибка добавления данных", err)
	// 		log.Fatal(err)
	// 	}
	//Инициализация тг бота
	bot, err := tgbotapi.NewBotAPI("8058580984:AAHIb2KHsG9msPViI7fwtOLNEKa_RztsY3U")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Приветствую, данный бот может показать тебе персонажей для взводов на Восход империи.")
				bot.Send(msg)
			case "help":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я могу: /n тебя зарегистрировать по команде /n начать работу по команде /start")
				bot.Send(msg)
			case "registration":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Для регистрации отравь свой игровой ник, игровой ID. Пример: DartV, 123456789")
				bot.Send(msg)
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I don't know that command")
				bot.Send(msg)
			}
		} else {

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			bot.Send(msg)
		}
	}

}
