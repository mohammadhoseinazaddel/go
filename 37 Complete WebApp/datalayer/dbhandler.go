package datalayer

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type SQLHandler struct {
	db *sql.DB
}

func CreateDBConnection(connString string) (*SQLHandler, error) {
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	return &SQLHandler{
		db: db,
	}, nil
}
