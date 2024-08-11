package repository

import (
	"context"

	"github.com/chanudinho/go-api-example/pkg/task/models"
	"gorm.io/gorm"
)

type taskRepository struct {
	*gorm.DB
}

func New(db *gorm.DB) *taskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) Create(ctx context.Context, task models.Task) (models.Task, error) {
	err := r.DB.
		WithContext(ctx).
		Model(&models.Task{}).
		Create(&task).
		Error

	return task, err
}

func (r *taskRepository) FindAll(ctx context.Context, task models.Task) ([]models.Task, error) {
	var tasks []models.Task
	err := r.DB.
		WithContext(ctx).
		Model(&models.Task{}).
		Where(task).
		Find(&tasks).
		Error

	return tasks, err
}

func (r *taskRepository) FindByID(ctx context.Context, id uint) (models.Task, error) {
	var t models.Task
	err := r.DB.
		WithContext(ctx).
		Model(&models.Task{}).
		Where("id = ?", id).
		First(&t).
		Error

	return t, err
}

func (r *taskRepository) Update(ctx context.Context, task models.Task) (models.Task, error) {
	err := r.DB.
		WithContext(ctx).
		Model(&models.Task{}).
		Where("id = ?", task.ID).
		Updates(&task).
		Error

	return task, err
}
