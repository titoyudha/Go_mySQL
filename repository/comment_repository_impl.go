package repository

import (
	"context"
	"database/sql"
	"go_mysql/entity"
)

type commentRepositoryImpl struct {
	DB *sql.DB
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
	panic("Find Here")
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	panic("Find Here")
}

func (repository *commentRepositoryImpl) DeleteById(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	panic("Delete")
}

func (repository *commentRepositoryImpl) DeleteAll(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	panic("Delete All")
}
