package db

import (
	"testing"
)

func TestNewMySQLConnection(t *testing.T) {
	cfg := DBConfig{
		User:     "testuser",
		Password: "testpass",
		Host:     "localhost",
		Port:     "3306",
		DBName:   "testdb",
	}

	db, err := NewMySQLConnection(cfg)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}
}
