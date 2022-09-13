package learngodatabase

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// do transaction
	for i := 0; i < 10000; i++ {
		email := "Punch" + strconv.Itoa(i) + "@rasyidev.id"
		comment := "Comment ke-" + strconv.Itoa(i)
		sqlQuery := "INSERT INTO comments(email,comment) VALUES (?,?)"

		_, err := tx.ExecContext(ctx, sqlQuery, email, comment)
		if err != nil {
			panic(err)
		}

		// rowsAffected, _ := res.RowsAffected()
		// lastId, _ := res.LastInsertId()
		// fmt.Println("Rows Affected\t:", rowsAffected)
		// fmt.Println("Last ID\t:", lastId)
	}
	tx.Commit()
	fmt.Println("berhasil mengeksekusi query")

}

/*
=== RUN   TestTransaction
berhasil mengeksekusi query
--- PASS: TestTransaction (1.25s)
PASS
ok      learn-go-database       1.294s
*/
