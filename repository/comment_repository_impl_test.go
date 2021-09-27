package repository

import (
	"context"
	"fmt"
	gomysql "go_mysql"
	"go_mysql/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(gomysql.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:    "repo@test.com",
		Comments: "Test Repo",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
