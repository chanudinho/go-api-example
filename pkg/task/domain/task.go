package domain

import (
	"time"

	"github.com/chanudinho/go-api-example/pkg/task/models"
)

type FindAllRequest struct {
	ID     uint   `json:"task_id" form:"task_id"`
	UserID uint   `json:"user_id" form:"user_id"`
	Title  string `json:"title" form:"title"`
}

type UpdateRequest struct {
	ID          uint       `json:"id" form:"id"`
	UserID      uint       `json:"user_id"`
	Title       string     `json:"title"`
	Summary     string     `json:"summary"`
	CompletedAt *time.Time `json:"completed_at"`
}

type Task struct {
	ID          uint       `json:"id"`
	UserID      uint       `json:"user_id"`
	Title       string     `json:"title"`
	Summary     string     `json:"summary"`
	CompletedAt *time.Time `json:"completed_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func (t *Task) FromModel(m models.Task) {
	t.ID = m.ID
	t.UserID = m.UserID
	t.Title = m.Title
	t.Summary = m.Summary
	t.CompletedAt = m.CompletedAt
	t.CreatedAt = m.CreatedAt
	t.UpdatedAt = m.UpdatedAt
}

func (t *Task) ToModel() models.Task {
	return models.Task{
		Title:       t.Title,
		UserID:      t.UserID,
		Summary:     t.Summary,
		CompletedAt: t.CompletedAt,
	}
}

func FromModels(tasks []models.Task) []Task {
	mapped := make([]Task, len(tasks))
	for i := 0; i < len(tasks); i++ {
		mapped[i].FromModel(tasks[i])
	}
	return mapped
}
