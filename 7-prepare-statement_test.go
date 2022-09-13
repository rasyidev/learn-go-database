package learngodatabase

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	sqlQuery := "INSERT INTO comments(email, comment) VALUES(?,?)"
	statement, err := db.PrepareContext(ctx, sqlQuery)
	if err != nil {
		panic(err.Error())
	}
	defer statement.Close()

	// simulasi batch insert
	for i := 0; i < 10000; i++ {
		email := "Punch" + strconv.Itoa(i) + "@rasyidev.id"
		comment := "Comment ke-" + strconv.Itoa(i)

		_, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err.Error())
		}

		// rowsAffected, _ := res.RowsAffected()
		// lastId, _ := res.LastInsertId()

		// fmt.Println("Rows Affected\t:", rowsAffected)
		// fmt.Println("Last ID\t:", lastId)
	}
	fmt.Println("Selesai mengeksekusi query")

}

/*
=== RUN   TestPrepareStatement
Selesai mengeksekusi query
--- PASS: TestPrepareStatement (5.99s)
PASS
ok      learn-go-database       6.042s
*/

func TestWithoutPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	querySQL := "INSERT INTO comments(email, comment) VALUES (?,?)"
	statement, err := db.PrepareContext(ctx, querySQL)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10000; i++ {
		email := "Punch" + strconv.Itoa(i) + "@rasyidev.id"
		comment := "Comment ke-" + strconv.Itoa(i)

		_, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		// rowsAffected, _ := res.RowsAffected()
		// lastId, _ := res.LastInsertId()
		// fmt.Println("Rows Affected\t:", rowsAffected)
		// fmt.Println("Last ID\t:", lastId)
	}
	fmt.Println("Selesai mengeksekusi query")

}

/*
=== RUN   TestWithoutPrepareStatement
Selesai mengeksekusi query
--- PASS: TestWithoutPrepareStatement (6.13s)
PASS
ok      learn-go-database       6.178s
*/
