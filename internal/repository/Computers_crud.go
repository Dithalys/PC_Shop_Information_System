package repository

import (
	"PC_Shop/internal/database"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Add_computer(db *sql.DB, title string, description string, price float32, user_id int) {
	var product_id int64
	query := "INSERT INTO Computers(title, overview, price, User_id) VALUES($1, $2, $3, $4) RETURNING id"
	err := db.QueryRow(query, title, description, price, user_id).Scan(&product_id)
	if err != nil {
		log.Println("Ошибка при добавлении компьютера:", err)
	} else {
		log.Printf("Добавлен компьютер [%d]\n", product_id)
	}

}

func Delete_computer(db *sql.DB, id int) {
	_, err := db.Exec("DELETE FROM Computers WHERE id = $1", id)
	if err != nil {
		log.Println("Ошибка при удалении компьютера:", err)
	} else {
		log.Printf("Удаление компьютера [%d]\n", id)
	}
}

func Update_comuter(db *sql.DB, title string, description string, price float32, id int) {
	_, err := db.Exec("UPDATE Computers set title = $1, overview = $2, price= $3 WHERE id = $4", title, description, price, id)
	if err != nil {
		log.Println("Ошибка при обновлении компьютера:", err)
	} else {
		log.Printf("Обновление компьютера [%d]\n", id)
	}
}

func Select_computers(db *sql.DB, user_id int) {
	rows, err := db.Query("SELECT * FROM Computers WHERE User_id = $1", user_id)
	if err != nil {
		log.Println("Ошибка при отрисовки компьютеров:", err)
	} else {
		log.Printf("Отрисовка списка компьютеров")
	}
	defer rows.Close()

	computers := []database.Computer{}
	for rows.Next() {
		computer := database.Computer{}
		err := rows.Scan(&computer.Id, &computer.Title, &computer.Overview, &computer.Price, &computer.User_id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		computers = append(computers, computer)
	}
	for _, c := range computers {
		fmt.Println(c.Id, c.Title, c.Overview, c.Price, c.User_id)
	}
}

// сделать возврат результата функциям
