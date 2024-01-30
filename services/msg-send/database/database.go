package database

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type conn struct {
	db *sql.DB
}

var instance *conn = nil

func InitDb() {
	var err error

	instance = &conn{}
	instance.db, err = sql.Open("pgx", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal(err)
	}
	err = instance.db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func GetInstance() *conn {
	if instance == nil {
		InitDb()
	}
	return instance
}

func WriteMessage(receiverId int64, senderUser string, body string) {
	tx, err := GetInstance().db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	rows, err := tx.Query("SELECT id FROM users WHERE username = ($1)", senderUser)
	if err != nil {
		log.Fatal(err)
	}

	var userId int64
	err = rows.Scan(&userId)
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO messages (sender_id, receiver_id, body) VALUES ($1, $2, $3)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(userId, receiverId, body)
	if err != nil {
		log.Fatal(err)
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}
}
