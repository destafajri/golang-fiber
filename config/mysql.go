package config

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/destafajri/golang-fiber/exception"
	_ "github.com/lib/pq"
)

func NewMySQLDatabase(configuration Config) *sql.DB {
	_, cancel := NewMySQLContext()
	defer cancel()

	database, err := sql.Open("mysql", configuration.Get("MYSQL_URL"))
	if err != nil {
		panic(err)
	}

	exception.PanicIfNeeded(err)

	err = database.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("MySQL Successfully connected!")

	return database
}

func NewMySQLContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}