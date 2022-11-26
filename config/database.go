package config

import (
	"database/sql"
	"os"
	"time"

	// database
	"github.com/go-graphql/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// Database - .
type Database struct {
	DB *sql.DB
}

// InitDB - .
func InitDB() *Database {
	db, err := sql.Open("mysql", os.Getenv("DATA_SOURCE"))
	if err != nil {
		logger.Log.Println(errors.Wrap(err, "database connection"))
	}

	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(time.Minute * 10)

	logger.Log.Println("database connected...")
	return &Database{DB: db}
}

// Close - close db
func (db *Database) Close() {
	err := db.DB.Close()
	if err != nil {
		logger.Log.Println(errors.Wrap(err, "database close"))
	}
}

// QueryRow -
func (db *Database) QueryRow(query string, args ...interface{}) *sql.Row {
	return db.DB.QueryRow(query, args...)
}

// Query -
func (db *Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.DB.Query(query, args...)
}

// Exec -
func (db *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.DB.Exec(query, args...)
}
