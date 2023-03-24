package test

import (
	"database/sql"
	"ryan-test-backend/helper"
	"time"
)

func SetupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:rootadmin@tcp(localhost:3306)/testing_ryan_test_backend")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func TruncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE products")
}
