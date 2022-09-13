package learngodatabase

import (
	"context"
	"fmt"
	"testing"
)

func TestSQLInjectionSafe(t *testing.T) {
	db := GetConnection()

	ctx := context.Background()

	username := "admin'; #"
	password := "ngawur"
	fmt.Println("username\t:", username)
	fmt.Println("password\t:", password)
	sqlQuery := "SELECT username, password FROM users WHERE username=? AND password=? LIMIT 1"
	fmt.Println(sqlQuery)

	rows, err := db.QueryContext(ctx, sqlQuery, username, password)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	if rows.Next() {
		fmt.Println("Berhasil Login")
	} else {
		fmt.Println("Gagal Login, username atau password salah")
	}
}

/*
$ go test -v -run TestSQLInjectionSafe
=== RUN   TestSQLInjectionSafe
username        : admin'; #
password        : ngawur
SELECT username, password FROM users WHERE username=? AND password=? LIMIT 1
Gagal Login, username atau password salah
--- PASS: TestSQLInjectionSafe (0.01s)
PASS
ok      learn-go-database       0.061s
*/
