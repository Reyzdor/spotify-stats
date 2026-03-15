package db

import (
	"database/sql"
	"fmt"
)

type DB struct {
	*sql.DB
}

func New(connStr string) (*DB, error) {
	sqlDB, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("Open DB: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("Ping DB: %w", err)
	}

	return &DB{sqlDB}, nil
}
