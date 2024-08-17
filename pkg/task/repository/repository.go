package repository

import (
	"context"

	"github.com/chanudinho/go-api-example/pkg/task/models"
)

type TaskRepo interface {
	Create(ctx context.Context, task models.Task) (models.Task, error)
	FindAll(ctx context.Context, task models.Task) ([]models.Task, error)
	FindByID(ctx context.Context, id uint) (models.Task, error)
	Update(ctx context.Context, task models.Task) (models.Task, error)
	Delete(ctx context.Context, id uint) (int64, error)
}
