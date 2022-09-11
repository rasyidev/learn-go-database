package learngodatabase

import (
	"context"
	"fmt"
	"testing"
)

func TestPerintahSQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlQuery := "INSERT INTO customer(id,name) VALUES ('rasyidev', 'Rasyidev');"
	_, err := db.ExecContext(ctx, sqlQuery)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Berhasil insert data ke dalam database")

}

func TestPerintahSelect(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	sqlQuery := "SELECT * FROM customer"
	rows, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(id, "|", name)
	}
}

/*
=== RUN   TestPerintahSelect
rasyidev | Rasyidev
taeri | Kim Taeri
yoona | Im Yoona
--- PASS: TestPerintahSelect (0.02s)
PASS
ok      learn-go-database       0.073s
*/
