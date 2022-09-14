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

	for i := 0; i < 10; i++ {
		comment := entity.Comment{
			Email:   "rasyidev" + strconv.Itoa(i) + "@test.id",
			Comment: "Comment ke-" + strconv.Itoa(i),
		}
		res, err := commentRepository.Insert(ctx, comment)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Berhasil memasukkan data:", res)
	}
}

/*=== RUN   TestInsertComment
Berhasil memasukkan data: {90039 rasyidev0@test.id Comment ke-0}
.....
Berhasil memasukkan data: {90040 rasyidev1@test.id Comment ke-9}
--- PASS: TestInsertComment (0.12s)
PASS
ok      learn-go-database/repository    0.168s*/

func TestFindByID(t *testing.T) {
	db := learngodatabase.GetConnection()
	ctx := context.Background()
	commentRepository := NewCommentRepository(db)

	res, err := commentRepository.FindById(ctx, 90039)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(res)
}

/*
$ go test -v -run TestFindByID
=== RUN   TestFindByID
{90039 rasyidev6@test.id Comment ke-6}
--- PASS: TestFindByID (0.02s)
PASS
ok      learn-go-database/repository    0.061s
*/

func TestFindAll(t *testing.T) {
	db := learngodatabase.GetConnection()
	ctx := context.Background()
	commentRepository := NewCommentRepository(db)

	res, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err.Error())
	}
	for _, each := range res {
		fmt.Println(each)
	}
}

/*
$ go test -v -run TestFindAll
=== RUN   TestFindAll
{90034 rasyidev1@test.id Comment ke-1}
{90035 rasyidev2@test.id Comment ke-2}
{90036 rasyidev3@test.id Comment ke-3}
{90037 rasyidev4@test.id Comment ke-4}
{90038 rasyidev5@test.id Comment ke-5}
{90039 rasyidev6@test.id Comment ke-6}
{90040 rasyidev7@test.id Comment ke-7}
{90041 rasyidev8@test.id Comment ke-8}
{90042 rasyidev9@test.id Comment ke-9}
--- PASS: TestFindAll (0.01s)
PASS
ok      learn-go-database/repository    0.063s
*/
