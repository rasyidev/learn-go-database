package learngodatabase

import (
	"context"
	"fmt"
	"testing"
)

func TestExecDenganParameter(t *testing.T) {
	db := GetConnection()
	username := "imyoona"
	password := "imyoonapass"

	ctx := context.Background()
	sqlQuery := "INSERT INTO users(username, password) VALUES(?, ?)"

	res, err := db.ExecContext(ctx, sqlQuery, username, password)
	if err != nil {
		panic(err.Error())
	}

	lastID, _ := res.LastInsertId()
	rowsAffected, _ := res.RowsAffected()

	fmt.Println("Last ID\t:", lastID)
	fmt.Println("rowsAffected\t:", rowsAffected)

	fmt.Println("Berhasil menambahkan user baru")

}

/*
$ go test -v -run TestExecDenganParameter
=== RUN   TestExecDenganParameter
Last ID : 0
rowsAffected    : 1
Berhasil menambahkan user baru
--- PASS: TestExecDenganParameter (0.02s)
PASS
ok      learn-go-database       0.059s
*/
