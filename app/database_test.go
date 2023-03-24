package app

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestNewDB(t *testing.T) {
	db := NewDB()

	// Test that the db is not nil
	if db == nil {
		t.Errorf("NewDB() returned nil, expected non-nil sql.DB instance")
	}

	// Test that the db has correct maximum open connections
	maxOpenConns := db.Stats().MaxOpenConnections
	if maxOpenConns != 20 {
		t.Errorf("NewDB() returned a db with maxOpenConns=%d, expected 20", maxOpenConns)
	}
}
