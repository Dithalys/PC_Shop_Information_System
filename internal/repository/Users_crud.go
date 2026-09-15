package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Add_user(name string, password string) {
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
	defer db.Close()

	result, err := db.Exec("INSERT INTO Users(username, password_hash) VALUES($1, $2)", name, password)
	if err != nil {
		panic(err)
	}
	fmt.Print(result.RowsAffected())
}

func Delete_user(id int) {
	connectionStr := "host=%s port=%s user=%s password=%s dbname=%s  sslmode=disable"
	connectionStr = fmt.Sprintf(connectionStr, os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM Users WHERE id = $1", id)
	if err != nil {
		panic(err)
	}
	fmt.Print(result.RowsAffected())
}

func Update_user(name string, password string, id int) {
	connectionStr := "host=%s port=%s user=%s password=%s dbname=%s  sslmode=disable"
	connectionStr = fmt.Sprintf(connectionStr, os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("UPDATE Users set username = $1, password_hash = $2 WHERE id = $3", name, password, id)
	if err != nil {
		panic(err)
	}
	fmt.Print(result.RowsAffected())
}
