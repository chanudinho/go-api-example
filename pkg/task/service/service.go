package service

import (
	"context"

	"github.com/chanudinho/go-api-example/pkg/task/domain"
	"github.com/chanudinho/go-api-example/pkg/task/repository"
)

type Service interface {
	Create(ctx context.Context, task domain.Task) (domain.Task, error)
	FindAll(ctx context.Context, task domain.FindAllRequest) ([]domain.Task, error)
	Update(ctx context.Context, task domain.UpdateRequest) (domain.Task, error)
}

type service struct {
	taskRepo repository.TaskRepo
}

func New(taskRepo repository.TaskRepo) Service {
	return &service{
		taskRepo: taskRepo,
	}
}
