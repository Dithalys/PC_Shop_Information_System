package main

import (
	"PC_Shop/internal/repository"
)

func main() {
	//epository.Add_user("user7", "password7")
	repository.Select_all_users()
	repository.Delete_user(9)
	repository.Update_user("user2", "password2", 7)
	repository.Select_all_users()
}
