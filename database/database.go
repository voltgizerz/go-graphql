package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	// database
	"github.com/go-graphql/pkg/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// Database - .
type Database struct {
	DB *sql.DB
}

// InitDB - .
func InitDB() (*Database, error) {
	dataSource := os.Getenv("DATA_SOURCE")
	if dataSource == "" {
		return nil, fmt.Errorf("DATA_SOURCE environment variable is empty or not set")
	}

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(time.Minute * 10)

	fmt.Println("Database connected...")
	return &Database{DB: db}, nil
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
