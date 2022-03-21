package dbrepo

import (
	"database/sql"

	"github.com/ImWhiteDevil/Bookings/internal/config"
	"github.com/ImWhiteDevil/Bookings/internal/repository"
)

type PostgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &PostgresDBRepo{
		App: a,
		DB:  conn,
	}

}