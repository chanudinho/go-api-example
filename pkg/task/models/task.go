package models

import (
	"time"

	"github.com/chanudinho/go-api-example/pkg/user/models"
	"gorm.io/gorm"
)

type Task struct {
	ID          uint           `gorm:"column:id;primaryKey;autoIncrement"`
	Title       string         `gorm:"column:title"`
	Summary     string         `gorm:"column:summary"`
	CompletedAt *time.Time     `gorm:"column:completed_at"`
	UserID      uint           `gorm:"column:user_id"`
	User        models.User    `gorm:"foreignKey:ID;references:UserID"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Task) TableName() string {
	return "tasks"
}
