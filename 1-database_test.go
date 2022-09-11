package learngodatabase

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql" // _ menjalankan init driver mysql
)

func TestConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/learn_go_database")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}
