package service

import (
	"context"
	"errors"

	"github.com/chanudinho/go-api-example/pkg/task/domain"
	"gorm.io/gorm"
)

func (s *service) Update(ctx context.Context, taskToUpdate domain.UpdateRequest) (domain.Task, error) {
	task, err := s.taskRepo.FindByID(ctx, taskToUpdate.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Task{}, errors.New("task not found")
		}
		return domain.Task{}, err
	}

	task.Title = taskToUpdate.Title
	task.Summary = taskToUpdate.Summary
	task.CompletedAt = taskToUpdate.CompletedAt

	updatedTask, err := s.taskRepo.Update(ctx, task)
	if err != nil {
		return domain.Task{}, err
	}

	var taskDomain domain.Task
	taskDomain.FromModel(updatedTask)
	return taskDomain, nil
}
