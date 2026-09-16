package repository

import (
	"PC_Shop/internal/database"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Add_user(db *sql.DB, name string, password string) {
	var user_id int64
	query := "INSERT INTO Users(username, password_hash) VALUES($1, $2) RETURNING id"
	err := db.QueryRow(query, name, password).Scan(&user_id)
	if err != nil {
		log.Println("Ошибка при добавлении поьлзователя:", err)
	} else {
		log.Printf("Добавлен пользователь [%d]\n", user_id)
	}

}

func Delete_user(db *sql.DB, id int) {
	_, err := db.Exec("DELETE FROM Users WHERE id = $1", id)
	if err != nil {
		log.Println("Ошибка при удалении поьлзователя:", err)
	} else {
		log.Printf("Удаление пользователя [%d]\n", id)
	}
}

func Update_user(db *sql.DB, name string, password string, id int) {
	_, err := db.Exec("UPDATE Users set username = $1, password_hash = $2 WHERE id = $3", name, password, id)
	if err != nil {
		log.Println("Ошибка при обновлении поьлзователя:", err)
	} else {
		log.Printf("Обновление пользователя [%d]\n", id)
	}
}

func Select_all_users(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM Users")
	if err != nil {
		log.Println("Ошибка при отрисовки поьлзователей:", err)
	} else {
		log.Printf("Отрисовка списка пользователей")
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

// сделать возврат результата функциям
