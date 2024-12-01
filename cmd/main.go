package main

import (
	"database/sql"
	"log"
	"trackify/cmd/api"
	"trackify/db"
	// gomail "github.com/go-mail/mail"
	_ "github.com/mattn/go-sqlite3"
)

// func sendMail(w http.ResponseWriter, r *http.Request) {
// 	msg := gomail.NewMessage()
//
// 	msg.SetHeader("From", "hello@demomailtrap.com")
// 	msg.SetHeader("To", "ignas321@gmail.com")
// 	msg.SetHeader("Subject", "Hello from the Mailtrap team")
//
// 	msg.SetBody("text/plain", "This is testing email")
//
// 	dialer := gomail.NewDialer("live.smtp.mailtrap.io", 587, "smtp@mailtrap.io", "c5be8bb79281771ab370033a0c50d291")
//
// 	if err := dialer.DialAndSend(msg); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
//
// 	io.WriteString(w, "Email sent")
// }

func main() {
	db, err := db.NewSQLiteStorage()
	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)

	api := api.NewServer(":8080", db)
	if err := api.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB successfully connected")
}
