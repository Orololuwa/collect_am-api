package models

import "time"

// Users is the user's model
type User struct {
	ID int `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
	Phone string `db:"phone"`
	Password string
	Avatar string
	CreatedAt time.Time
	UpdatedAt time.Time
	Business Business
	// One to One
}