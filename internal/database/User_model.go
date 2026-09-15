package database

type User struct {
	Id            int
	Username      string
	Password_hash string
	Balance       float32
}
