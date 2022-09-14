package repository

import (
	"context"
	"fmt"
	learngodatabase "learn-go-database"
	"learn-go-database/entity"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInsertComment(t *testing.T) {
	db := learngodatabase.GetConnection()
	ctx := context.Background()
	commentRepository := NewCommentRepository(db)

	for i := 0; i < 2; i++ {
		comment := entity.Comment{
			Email:   "rasyidev" + strconv.Itoa(i) + "@test.id",
			Comment: "Comment ke-" + strconv.Itoa(i),
		}
		res, err := commentRepository.Insert(ctx, comment)
		if err != nil {
			panic(err)
		}
		fmt.Println("Berhasil memasukkan data:", res)
	}
}
