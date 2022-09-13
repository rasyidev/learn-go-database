package learngodatabase

import (
	"context"
	"fmt"
	"testing"
)

func TestSQLInjection(t *testing.T) {
	db := GetConnection()

	ctx := context.Background()

	username := "admin'; #"
	password := "ngawur"
	sqlQuery := "SELECT username, password FROM users WHERE username='" + username +
		"' AND password='" + password + "' LIMIT 1"
	fmt.Println(sqlQuery)

	rows, err := db.QueryContext(ctx, sqlQuery)
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
$ go test -v -run TestSQLInjection
username: admin
password: admin
=== RUN   TestSQLInjection
SELECT username, password FROM users WHERE username='admin'; #' AND password='ngawur' LIMIT 1
Berhasil Login
--- PASS: TestSQLInjection (0.02s)
PASS
ok      learn-go-database       0.074s
*/
