package config

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/destafajri/golang-fiber/exception"
	_ "github.com/lib/pq"

	// "github.com/golang-migrate/migrate/v4"
    // "github.com/golang-migrate/migrate/v4/database/postgres"
    // _ "github.com/golang-migrate/migrate/v4/source/file"
)

func NewPostgreDatabase(configuration Config) *sql.DB {
	_, cancel := NewPostgresContext()
	defer cancel()

	database, err := sql.Open("postgres", configuration.Get("POSTGRES_URL"))
	exception.PanicIfNeeded(err)
	defer database.Close()

	err = database.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Postgres Successfully connected!")

	// driver, err := postgres.WithInstance(database, &postgres.Config{})
	// exception.PanicIfNeeded(err)

    // m, err := migrate.NewWithDatabaseInstance(
    //     "file://migrations/postgres",
    //     "postgres", driver)
	// exception.PanicIfNeeded(err)

    // m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run

	// log.Println("Migration Postgres Successfully!")

	return database
}

func NewPostgresContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
