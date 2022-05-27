package repository

import (
	"context"
	"go-advanced-rest-websockets/models"
)

func InsertPost(ctx context.Context, post *models.Post) error {
	return implementation.InsertPost(ctx, post)
}

func GetPostById(ctx context.Context, id string) (*models.Post, error) {
	return implementation.GetPostById(ctx, id)
}

func UpdatePost(ctx context.Context, post *models.Post) error {
	return implementation.UpdatePost(ctx, post)
}

func DeletePostById(ctx context.Context, id string, userId string) error {
	return implementation.DeletePostById(ctx, id, userId)
}

func ListPost(ctx context.Context, limit uint64, page uint64) ([]*models.Post, error) {
	return implementation.ListPost(ctx, limit, page)
}
