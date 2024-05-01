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
}

type Room struct {
	ID int `json:"id"`
	RoomName string `json:"roomName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

}

type Restriction struct {
	ID int
	RestrictionName string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Reservation struct {
	ID int
	FirstName string
	LastName  string
	Email     string
	Phone string
	StartDate time.Time
	EndDate time.Time
	RoomID int
	CreatedAt time.Time
	UpdatedAt time.Time
	Room Room
}

type RoomRestriction struct {
	ID int
	StartDate time.Time
	EndDate time.Time
	RoomID int
	ReservationID int
	RestrictionID int
	CreatedAt time.Time
	UpdatedAt time.Time
	Room Room
	Reservation Reservation
	Restriction Restriction
}
