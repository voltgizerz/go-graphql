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
