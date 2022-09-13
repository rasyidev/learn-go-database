package learngodatabase

import (
	"context"
	"fmt"
	"testing"
	"time"
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

func TestPerintahSelectComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	sqlQuery := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, email string
		var balance int32
		var rating float32
		var birth_date, created_at time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &created_at)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("----------------------------------------------------------------------------")
		fmt.Println("ID\t\t:", id)
		fmt.Println("Name\t\t:", name)
		fmt.Println("Email\t\t:", email)
		fmt.Println("Balance\t\t:", balance)
		fmt.Println("Rating\t\t:", rating)
		fmt.Println("Birth Date\t:", birth_date)
		fmt.Println("Maried\t\t:", married)
		fmt.Println("Created At\t:", created_at)
	}
}

/*
$ go test -v -run TestPerintahSelectComplex
=== RUN   TestPerintahSelectComplex
--- FAIL: TestPerintahSelectComplex (0.01s)
panic: sql: Scan error on column index 5, name "birth_date": unsupported Scan, storing driver.Value type []uint8 into typ

Driver MySQL untuk golang:
DATE, DATETIME, TIMESTAMP -> []byte -> parsing ke string -> time.Time (tapi ribet)
setting aja di driver golang parseDate=True

*/

/*
$ go test -v -run TestPerintahSelectComplex
=== RUN   TestPerintahSelectComplex
----------------------------------------------------------------------------
ID              : 1
Name            : Rasyidev
Email           : ceo@rasyidev.com
Balance         : 23
Rating          : 4.8
Birth Date      : 2002-08-23 00:00:00 +0000 UTC
Maried          : false
Created At      : 2022-09-12 22:49:30 +0000 UTC
----------------------------------------------------------------------------
ID              : 2
Name            : Kim Taeri
Email           : taeri@rasyidev.com
Balance         : 43
Rating          : 4.5
Birth Date      : 1993-03-14 00:00:00 +0000 UTC
Maried          : false
Created At      : 2022-09-12 22:49:30 +0000 UTC
----------------------------------------------------------------------------
ID              : 3
Name            : Im Yoona
Email				    : yoona@rasyidev.com
Balance         : 14
Rating          : 4.3
Birth Date      : 1995-07-23 00:00:00 +0000 UTC
Maried          : false
Created At      : 2022-09-12 22:49:30 +0000 UTC
--- PASS: TestPerintahSelectComplex (0.02s)
PASS
ok      learn-go-database       0.073s
*/
