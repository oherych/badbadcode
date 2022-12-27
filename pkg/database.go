package pkg

import (
	"database/sql"
	"os"
)

var (
	DB *sql.DB
)

func init() {
	var err error
	DB, err = sql.Open("driver-name", os.Getenv("DB_DSN"))
	if err != nil {
		panic(err)
	}
}
