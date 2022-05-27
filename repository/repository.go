package repository

import (
	"context"
	"go-advanced-rest-websockets/models"
)

type Repository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)

	InsertPost(ctx context.Context, post *models.Post) error
	GetPostById(ctx context.Context, id string) (*models.Post, error)
	UpdatePost(ctx context.Context, post *models.Post) error

	Close() error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func Close() error {
	return implementation.Close()
}
