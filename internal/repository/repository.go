package repository

import "github.com/ImWhiteDevil/Booking/internal/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) error
}
