package repository

import (
	"PC_Shop/internal/database"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Add_computer_component(db *sql.DB, computer_id int, component_id int, quaintity int, user_id int) {
	var product_id int64
	query := "INSERT INTO Computer_Components(Computer_id, Component_id, quantity, User_id) VALUES($1, $2, $3, $4) RETURNING id"
	err := db.QueryRow(query, computer_id, component_id, quaintity, user_id).Scan(&product_id)
	if err != nil {
		log.Println("Ошибка при добавлении компонена:", err)
	} else {
		log.Printf("Добавлен компонент [%d]\n", product_id)
	}

}

func Delete_computer_component(db *sql.DB, id int) {
	_, err := db.Exec("DELETE FROM Computer_Components WHERE id = $1", id)
	if err != nil {
		log.Println("Ошибка при удалении компонента:", err)
	} else {
		log.Printf("Удаление компонента [%d]\n", id)
	}
}

func Update_computer_component(db *sql.DB, computer_id int, component_id int, quaintity int, id int) {
	_, err := db.Exec("UPDATE Computer_Components set Computer_id = $1, Component_id = $2, quantity = $3 WHERE id = $4", computer_id, component_id, quaintity, id)
	if err != nil {
		log.Println("Ошибка при обновлении компонента:", err)
	} else {
		log.Printf("Обновление компонента [%d]\n", id)
	}
}

func Select_computer_components(db *sql.DB, user_id int, computer_id int) {
	rows, err := db.Query("SELECT * FROM Computer_Components WHERE User_id = $1 and Computer_id = $2", user_id, computer_id)
	if err != nil {
		log.Println("Ошибка при отрисовки компонетов:", err)
	} else {
		log.Printf("Отрисовка списка компонентов")
	}
	defer rows.Close()

	components := []database.Copmuter_components{}
	for rows.Next() {
		component := database.Copmuter_components{}
		err := rows.Scan(&component.Id, &component.Computer_id, &component.Component_id, &component.Quantity, &component.User_id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		components = append(components, component)
	}
	for _, c := range components {
		fmt.Println(c.Id, c.Computer_id, c.Component_id, c.Quantity, c.User_id)
	}
}

// сделать возврат результата функциям
