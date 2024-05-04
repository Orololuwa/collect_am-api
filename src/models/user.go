package models

import "time"

// Users is the user's model
type User struct {
	ID int `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
	Phone string `db:"phone"`
	Password string `db:"password"`
	Avatar string `db:"avatar"`
	Gender string `db:"gender"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Business Business
	// One to One
}