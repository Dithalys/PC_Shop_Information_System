package repository

import (
	"PC_Shop/internal/database"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Add_component(db *sql.DB, title string, description string, price float32, quaintity int, user_id int) {
	var product_id int64
	query := "INSERT INTO Components(title, overview, price, quantity, User_id) VALUES($1, $2, $3, $4, $5) RETURNING id"
	err := db.QueryRow(query, title, description, price, quaintity, user_id).Scan(&product_id)
	if err != nil {
		log.Println("Ошибка при добавлении компонена:", err)
	} else {
		log.Printf("Добавлен компонент [%d]\n", product_id)
	}

}

func Delete_component(db *sql.DB, id int) {
	_, err := db.Exec("DELETE FROM Components WHERE id = $1", id)
	if err != nil {
		log.Println("Ошибка при удалении компонента:", err)
	} else {
		log.Printf("Удаление компонента [%d]\n", id)
	}
}

func Update_component(db *sql.DB, title string, description string, price float32, quaintity int, id int) {
	_, err := db.Exec("UPDATE Components set title = $1, overview = $2, price= $3, quantity = $4 WHERE id = $5", title, description, price, quaintity, id)
	if err != nil {
		log.Println("Ошибка при обновлении компонента:", err)
	} else {
		log.Printf("Обновление компонента [%d]\n", id)
	}
}

func Select_components(db *sql.DB, user_id int) {
	rows, err := db.Query("SELECT * FROM Components WHERE User_id = $1", user_id)
	if err != nil {
		log.Println("Ошибка при отрисовки компонетов:", err)
	} else {
		log.Printf("Отрисовка списка компонентов")
	}
	defer rows.Close()

	components := []database.Component{}
	for rows.Next() {
		component := database.Component{}
		err := rows.Scan(&component.Id, &component.Title, &component.Overview, &component.Price, &component.Quantity, &component.User_id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		components = append(components, component)
	}
	for _, c := range components {
		fmt.Println(c.Id, c.Title, c.Overview, c.Price, c.Quantity, c.User_id)
	}
}

// сделать возврат результата функциям
