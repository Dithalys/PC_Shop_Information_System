package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	connectionStr := "host=%s port=%s user=%s password=%s dbname=%s  sslmode=disable"
	connectionStr = fmt.Sprintf(connectionStr, os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}
	return db
}
