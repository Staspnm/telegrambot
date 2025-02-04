package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
)

type userBD struct {
	user_pk              float32
	user_uuid            string
	username             string
	role_id              float32
	usertg               string
	enabled              bool
	phone                string
	dop_info             string
	last_connection_time time.Time
	last_network_address string
	update_time          time.Time
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	host := "localhost"
	port := 5000
	user := "postgres"
	password := "fu2xo4q"
	dbname := "postgres"

	// строка подключения
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// открыть базу данных
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	// закрыть базу данных
	defer db.Close()
	// проверить подключение
	err = db.Ping()
	CheckError(err)
	fmt.Println("Connected!")

	//запрос в базу
	rows, err := db.Query("SELECT user_pk, user_uuid, username, role_id, usertg, enabled, phone, dop_info, last_connection_time, last_network_address, update_time FROM rote.users where username = 'DartV'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	usrs := make([]*userBD, 0)
	for rows.Next() {
		usr := new(userBD)
		err := rows.Scan(&usr.user_pk,
			&usr.user_uuid,
			&usr.username,
			&usr.role_id,
			&usr.usertg,
			&usr.enabled,
			&usr.phone,
			&usr.dop_info,
			&usr.last_connection_time,
			&usr.last_network_address,
			&usr.update_time)
		if err != nil {
			log.Fatal(err)
		}
		usrs = append(usrs, usr)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, usr := range usrs {
		fmt.Printf("%d, %s, %s, %d, %s, %s, %s, %s, %s, %s, %s", usr.user_pk, usr.user_uuid, usr.username, usr.role_id, usr.usertg, usr.enabled, usr.phone, usr.dop_info, usr.last_connection_time, usr.last_network_address, usr.update_time)
	}
	// конец запроса

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
