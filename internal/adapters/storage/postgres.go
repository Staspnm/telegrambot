package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type UserBD struct {
	User_pk              float32
	User_uuid            string
	Username             string
	Role_id              float32
	Usertg               string
	Enabled              bool
	Phone                string
	Dop_info             string
	Last_connection_time time.Time
	Last_network_address string
	Update_time          time.Time
}

type Storage struct {
	db *sql.DB
}

func New() *Storage {
	host := "localhost"
	port := 5012
	user := "postgres"
	password := "fu2xo4q"
	dbname := "postgres"

	// строка подключения
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// открыть базу данных
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	// закрыть базу данных
	// defer db.Close() // Здесь нельзя закрывать базу тк сразу на выходе из этой функции коннект к бд закроется
	// проверить подключение
	err = db.Ping()
	CheckError(err)
	fmt.Println("Connected!")

	return &Storage{
		db: db,
	}
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// запрос в базу
func (st *Storage) GetUser() ([]*UserBD, error) {
	rows, err := st.db.Query("SELECT user_pk, user_uuid, username, role_id, usertg, enabled, phone, dop_info, last_connection_time, last_network_address, update_time FROM rote.users where username = 'DartV'")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	usrs := make([]*UserBD, 0)
	for rows.Next() {
		usr := new(UserBD)
		err := rows.Scan(&usr.User_pk,
			&usr.User_uuid,
			&usr.Username,
			&usr.Role_id,
			&usr.Usertg,
			&usr.Enabled,
			&usr.Phone,
			&usr.Dop_info,
			&usr.Last_connection_time,
			&usr.Last_network_address,
			&usr.Update_time)
		if err != nil {
			log.Fatal(err)
		}
		usrs = append(usrs, usr)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return usrs, nil
}

func (st *Storage) Close() {
	st.db.Close()
}
