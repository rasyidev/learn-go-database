package learngodatabase

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:bismillah,.@tcp(localhost:3306)/learn_go_database?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	// Jangan di defer, nanti langsung ditutup,
	// defer db.Close()

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
