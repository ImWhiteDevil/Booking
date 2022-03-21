package drivers

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type DB struct {
	SQL *sql.DB
}

var dbconn = &DB{}

const maxOpendbconn = 10
const maxIdledbconn = 5
const maxDBLifeTime = 5 * time.Minute

func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDataBase(dsn)
	if err != nil {
		panic(err)
	}
	d.SetMaxOpenConns(maxOpendbconn)
	d.SetConnMaxIdleTime(maxIdledbconn)
	d.SetConnMaxLifetime(maxDBLifeTime)

	dbconn.SQL = d

	err = testDB(d)
	if err != nil {
		return nil, err
	}
	return dbconn, nil
}

func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}
	return nil

}

func NewDataBase(dsn string) (*sql.DB, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("unable to open database", err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
