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
	//repository.Add_component(db, "i5", "duiiajo", 23.7, 5, 1)
	//repository.Delete_component(db, 4)
	//repository.Select_components(db, 35)
	//repository.Update_component(db, "sdf", "", 23, 1, 2)

	// запросы для компьютеров
	//repository.Add_computer(db, "comp", "asdsadcklnascasnc", 234.43, 1)
	//repository.Delete_computer(db, 1)
	//repository.Update_comuter(db, "sadf", "", 21.2, 2)
	//repository.Select_computers(db, 35)

	// запросы для компонетов компьютера
	//repository.Add_computer_component(db, 2, 1, 1, 1)
	//repository.Update_computer_component(db, 1, 1, 2, 1)
	//repository.Delete_computer_component(db, 2)
	//repository.Select_computer_components(db, 1, 2)
}
