package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go-database/model"
)

type CommentRepository interface {
	GetAll(ctx context.Context) ([]model.Comment, error)
	GetByID(ctx context.Context, id int) (*model.Comment, error)
	Create(ctx context.Context, comment model.Comment) error
	Update(ctx context.Context, comment model.Comment) error
	Delete(ctx context.Context, id int) error
}

type commentRepositoryImpl struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) (CommentRepository, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	return &commentRepositoryImpl{
		db: db,
	}, nil
}

func (r *commentRepositoryImpl) GetAll(ctx context.Context) ([]model.Comment, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, email, comment FROM comments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []model.Comment
	for rows.Next() {
		var comment model.Comment
		if err := rows.Scan(&comment.ID, &comment.Email, &comment.Comment); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (r *commentRepositoryImpl) GetByID(ctx context.Context, id int) (*model.Comment, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, email, comment FROM comments WHERE id = ?", id)
	var comment model.Comment

	if err := row.Scan(&comment.ID, &comment.Email, &comment.Comment); err != nil {
		return nil, err
	}

	return &comment, nil
}

func (r *commentRepositoryImpl) Create(ctx context.Context, comment model.Comment) error {
	_, err := r.db.ExecContext(ctx, "insert into comments(email, comment) values (?, ?)", comment.Email, comment.Comment)
	return err
}

func (r *commentRepositoryImpl) Update(ctx context.Context, comment model.Comment) error {
	_, err := r.db.ExecContext(ctx, "update comments set email = ?, comment = ? where id = ?", comment.Email, comment.Comment, comment.ID)
	return err
}

func (r *commentRepositoryImpl) Delete(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, "delete from comments where id = ?", id)
	return err
}
