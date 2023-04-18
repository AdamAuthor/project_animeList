package models

type User struct {
	ID       int    `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Name     string `json:"name" db:"name"`
	Gender   string `json:"gender" db:"gender"`
	Age      int    `json:"age" db:"age"`
}
