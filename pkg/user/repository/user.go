package repository

import (
	"context"

	"github.com/chanudinho/go-api-example/pkg/user/models"
	"gorm.io/gorm"
)

type userRepository struct {
	*gorm.DB
}

func New(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, user models.User) (models.User, error) {
	err := r.DB.
		WithContext(ctx).
		Create(&user).
		Error

	return user, err
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User

	err := r.DB.
		WithContext(ctx).
		Where("email = ?", email).
		First(&user).
		Error

	return user, err
}
