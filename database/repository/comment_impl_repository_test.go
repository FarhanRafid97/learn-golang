package repository

import (
	"belajar-go-routine/database/entity"
	"context"
	"fmt"
	"testing"
)

func TestRepositoryInsert(t *testing.T) {
	commentRepositoy := NewCommentRepository(GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email:   "TestRepo@gmail.com",
		Comment: "Test repo",
	}
	result, err := commentRepositoy.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

}

func TestRepositoryFindById(t *testing.T) {
	commentRepositoy := NewCommentRepository(GetConnection())
	ctx := context.Background()

	result, err := commentRepositoy.FindById(ctx, 47)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

}
func TestRepositoryFindAll(t *testing.T) {
	commentRepositoy := NewCommentRepository(GetConnection())
	ctx := context.Background()

	comments, err := commentRepositoy.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}

}
