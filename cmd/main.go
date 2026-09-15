package main

import (
	"PC_Shop/internal/database"
	"PC_Shop/internal/repository"
)

func main() {
	db := database.Connect()
	defer db.Close()
	//repository.Add_user(db, "user1211", "password")
	//repository.Select_all_users(db)
	//repository.Delete_user(db, 98)
	//repository.Update_user(db, "user2", "password2", 267)
	repository.Select_all_users(db)
}
