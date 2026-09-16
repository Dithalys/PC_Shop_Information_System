package main

import (
	"PC_Shop/internal/database"
)

func main() {
	db := database.Connect()
	defer db.Close()

	// запросы для пользователей
	//repository.Add_user(db, "user1211", "password")
	//repository.Select_all_users(db)
	//repository.Delete_user(db, 98)
	//repository.Update_user(db, "user2", "password2", 267)

	// запросы для компонентов
	//repository.Add_component(db, "i5", "duiiajo", 23.7, 5, 35)
	//repository.Delete_component(db, 4)
	//repository.Select_components(db, 35)
	//repository.Update_component(db, "sdf", "", 23, 1, 2)
}
