package db

import (
	"database/sql"
	"log"
	"os"

	// database
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type Database struct {
	Con *sql.DB
}

func Connect() *Database {
	con, err := sql.Open("mysql", os.Getenv("DATA_SOURCE"))
	if err != nil {
		log.Println(errors.Wrap(err, "database connection"))
	}
	con.SetMaxOpenConns(10)

	log.Println("database connected...")
	return &Database{Con: con}
}

// Close - close db
func (db *Database) Close() {
	err := db.Con.Close()
	if err != nil {
		log.Println(errors.Wrap(err, "database close"))
	}
}

// QueryRow -
func (db *Database) QueryRow(query string, args ...interface{}) *sql.Row {
	log.Println(query, args)
	return db.Con.QueryRow(query, args...)
}

// Query -
func (db *Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	log.Println(query, args)
	return db.Con.Query(query, args...)
}

// Exec -
func (db *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	log.Println(query, args)
	return db.Con.Exec(query, args...)
}
