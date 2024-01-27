package database

import (
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

	// Insert test user to database
	stmt, err := GetInstance().db.Prepare("INSERT INTO users (username, password) VALUES ($1, $2)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(os.Getenv("SENDER_USER"), os.Getenv("SENDER_PASSWORD"))
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
