package dbrepo

import (
	"database/sql"

	"github.com/ImWhiteDevil/Booking/internal/config"
	"github.com/ImWhiteDevil/Booking/internal/repository"
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
