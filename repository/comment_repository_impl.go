package repository

import (
	"context"
	"database/sql"
	"errors"
	"learn-go-database/entity"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	sqlQuery := "INSERT INTO comments(email,comment) VALUES (?,?)"
	res, err := repository.DB.ExecContext(ctx, sqlQuery, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, Id int32) (entity.Comment, error) {
	sqlQuery := "SELECT id, email, comment FROM comment WHERE id=? LIMIT 1"
	row, err := repository.DB.QueryContext(ctx, sqlQuery, Id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer row.Close()

	if row.Next() {
		// ada
		row.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil

	} else {
		// tidak ada
		return comment, errors.New("Id " + strconv.Itoa(int(Id)) + " tidak ditemukan")
	}
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	sqlQuery := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, sqlQuery)
	comments := []entity.Comment{}
	if err != nil {
		return comments, err
	}
	defer rows.Close()

	if rows.Next() {
		for rows.Next() {
			comment := entity.Comment{}
			rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
			comments = append(comments, comment)
		}
	}
	return comments, nil
}
