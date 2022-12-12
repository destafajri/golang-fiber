package config

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/destafajri/golang-fiber/exception"
	_ "github.com/lib/pq"
)

func NewPostgreDatabase(configuration Config) *sql.DB {
	_, cancel := NewPostgresContext()
	defer cancel()

	database, err := sql.Open("postgres", configuration.Get("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	exception.PanicIfNeeded(err)

	err = database.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Postgres Successfully connected!")

	return database
}

func NewPostgresContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
