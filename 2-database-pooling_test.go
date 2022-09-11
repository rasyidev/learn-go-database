package learngodatabase

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/learn_go_database")
	if err != nil {
		panic(err.Error())
	}
	// Jangan di defer, nanti langsung ditutup, kecuali bikin WaitGroup
	// defer db.Close()

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
