package service

import (
	"context"

	"github.com/chanudinho/go-api-example/pkg/task/domain"
	"github.com/chanudinho/go-api-example/pkg/task/models"
)

func (s *service) FindAll(ctx context.Context, task domain.FindAllRequest) ([]domain.Task, error) {
	tasks, err := s.taskRepo.FindAll(ctx, models.Task{
		ID:     task.ID,
		UserID: task.UserID,
		Title:  task.Title,
	})
	if err != nil {
		return []domain.Task{}, err
	}

	return domain.FromModels(tasks), nil
}
