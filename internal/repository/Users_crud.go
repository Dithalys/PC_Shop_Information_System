package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Add_user(name string, password string) {
	connectionStr := "host=localhost user=%s password=%s dbname=%s port=5433 sslmode=disable"
	connectionStr = fmt.Sprintf(connectionStr, "User", "Maputa129", "PC_Shop_DB")
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
	connectionStr := "host=localhost user=%s password=%s dbname=%s port=5433 sslmode=disable"
	connectionStr = fmt.Sprintf(connectionStr, "User", "Maputa129", "PC_Shop_DB")
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
	connectionStr := "host=localhost user=%s password=%s dbname=%s port=5433 sslmode=disable"
	connectionStr = fmt.Sprintf(connectionStr, "User", "Maputa129", "PC_Shop_DB")
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
