package test

import (
	"context"
	"fmt"
	"go-database/config"
	"go-database/model"
	"go-database/repository"
	"testing"
	"time"
)

func TestRepositoryPattern(t *testing.T) {
	config.ConnectDB()

	repo, err := repository.NewCommentRepository(config.DB)
	if err != nil {
		t.Fatalf("Failed to create repository: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// insert comment
	if err := repo.Create(ctx, model.Comment{
		Email:   "test@example.com",
		Comment: "This is a test comment",
	}); err != nil {
		fmt.Println("Gagal insert:", err)
	}

	// get all
	comments, err := repo.GetAll(ctx)
	if err != nil {
		fmt.Println("Gagal ambil data:", err)
	}

	fmt.Println("Comments:", comments)
}
