package repositories

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

var (
	db    *sqlx.DB
	dbURI string
)

// Bootstrap configures MySQL connection
func Bootstrap() *sqlx.DB {
	dbURI = os.Getenv("MYSQL_URI")
	db = Connect()

	err := db.Ping()

	if err != nil {
		panic(fmt.Sprintf("failed to connect to database. error: %s", err))
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)

	return db
}

// Connect is a helper function to stablish a new MySQL connection
func Connect() *sqlx.DB {

	if db != nil {
		if err := db.Ping(); err != nil {
			db = sqlx.MustConnect("mysql", dbURI)
		}

		return db
	}

	db = sqlx.MustConnect("mysql", dbURI)

	return db
}
