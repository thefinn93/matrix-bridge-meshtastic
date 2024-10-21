package db

import (
	"database/sql"
	"embed"

	"git.janky.solutions/finn/matrix-meshtastic-bridge-go/config"
	_ "github.com/mattn/go-sqlite3"
	goose "github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
)

//go:embed migrations
var migrations embed.FS

func Migrate() error {
	logrus.Info("running database migrations")

	_, dbConn, err := Get()
	if err != nil {
		return err
	}
	defer dbConn.Close()

	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		return err
	}

	if err := goose.Up(dbConn, "migrations"); err != nil {
		return err
	}

	return nil
}

// Get the database and closable DB object
// example usage:
//
//	  queries, dbconn, err := db.Get()
//	  if err != nil {
//		   return err
//	  }
//	  defer dbconn.Close()
func Get() (*Queries, *sql.DB, error) {
	db, err := sql.Open("sqlite3", config.C.Database)
	if err != nil {
		return nil, nil, err
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return nil, nil, err
	}

	return New(db), db, nil
}

// NullableString accepts a string and returns an sql.NullString
// if the input string is zero-length, the output will be marked
// null
func NullableString(s string) sql.NullString {
	return sql.NullString{
		Valid:  s != "",
		String: s,
	}
}
