package main

import (
	"PC_Shop/internal/repository"
)

func main() {
	//repository.Add_user("user1", "password1")
	//repository.Delete_user(6)
	repository.Update_user("user2", "password2", 7)
}
