package util

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB(username string, password string, dbname string) (*sql.DB, error) {
	connStr := "user=" + username + " password=" + password + " dbname=" + dbname + " sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
