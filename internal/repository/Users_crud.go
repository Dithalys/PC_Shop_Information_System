package repository

import (
	"PC_Shop/internal/database"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Add_user(db *sql.DB, name string, password string) {
	result, err := db.Exec("INSERT INTO Users(username, password_hash) VALUES($1, $2)", name, password)
	if err != nil {
		panic(err)
	}
	fmt.Print(result.RowsAffected())
}

func Delete_user(db *sql.DB, id int) {
	result, err := db.Exec("DELETE FROM Users WHERE id = $1", id)
	if err != nil {
		panic(err)
	}
	fmt.Print(result.RowsAffected())
}

func Update_user(db *sql.DB, name string, password string, id int) {
	result, err := db.Exec("UPDATE Users set username = $1, password_hash = $2 WHERE id = $3", name, password, id)
	if err != nil {
		panic(err)
	}
	fmt.Print(result.RowsAffected())
}

func Select_all_users(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM Users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	users := []database.User{}

	for rows.Next() {
		user := database.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Password_hash, &user.Balance)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, user)
	}
	for _, u := range users {
		fmt.Println(u.Id, u.Username, u.Password_hash, u.Balance)
	}
}
