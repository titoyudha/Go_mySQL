package repository

import (
	"context"
	"database/sql"
	"errors"

	"go_mysql/entity"
	"strconv"
	//"golang.org/x/crypto/openpgp/errors"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comments(email, comment) VALUE(?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Comments)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, nil
	}
	comment.Id = int32(id)
	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		//jika ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Comments)
		return comment, nil
	} else {
		//Jika Tidak ada
		return comment, errors.New("Id" + strconv.Itoa(int(id)) + "Not Found")
	}
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		//jika ada
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comments)
		comments = append(comments, comment)
	}
	return comments, nil
}

func (repository *commentRepositoryImpl) DeleteById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "DELETE id, email, comment FROM comments WHERE ID = ? "
	rows, err := repository.DB.QueryContext(ctx, script, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		//jika ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Comments)
		return comment, nil
	} else {
		//Jika Tidak ada
		return comment, errors.New("Id" + strconv.Itoa(int(id)) + "Not Found")
	}
}

func (repository *commentRepositoryImpl) DeleteAll(ctx context.Context) ([]entity.Comment, error) {
	script := "DELETE id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {

		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comments)
		comments = append(comments, comment)
	}
	return comments, nil
}
