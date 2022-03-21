package models

import "time"

type User struct{
	ID int
	FirstName string
	LastName string
	Email string
	Password string
	AccessLevel int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Room struct{
	ID int
	RoomName string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Restriction struct{
	ID int
	ReestrictionName string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Reservation struct{
	ID int
	FirstName string
	LastName string
	Email string
	Phone string
	StartDate time.Time
	EndDate time.Time
	RoomID int
	CreatedAt time.Time
	UpdatedAt time.Time
	Room Room
}

type RoomRestriction struct{
	ID int
	StartDate time.Time
	EndDate time.Time
	RoomID int
	ReservationID int
	ReestrictionID int
	CreatedAt time.Time
	UpdatedAt time.Time
	Room Room
	Reservation Reservation	
	Restriction Restriction
}