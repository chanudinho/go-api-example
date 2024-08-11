package repository

import (
	"context"

	"github.com/chanudinho/go-api-example/pkg/user/models"
)

type UserRepo interface {
	Create(ctx context.Context, user models.User) (models.User, error)
	FindByEmail(ctx context.Context, email string) (models.User, error)
}
