package service

import (
	"context"

	"github.com/chanudinho/go-api-example/pkg/task/domain"
)

func (s *service) Create(ctx context.Context, taskToCreate domain.Task) (domain.Task, error) {
	task, err := s.taskRepo.Create(ctx, taskToCreate.ToModel())
	if err != nil {
		return domain.Task{}, err
	}

	var taskDomain domain.Task
	taskDomain.FromModel(task)
	return taskDomain, nil
}
