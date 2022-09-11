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
