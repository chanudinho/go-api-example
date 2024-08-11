package service

import (
	"context"

	"github.com/chanudinho/go-api-example/pkg/user/domain"
	"github.com/chanudinho/go-api-example/pkg/user/repository"
)

type Service interface {
	Create(ctx context.Context, user domain.User) (domain.User, error)
	Login(ctx context.Context, user domain.User) (string, error)
}

type service struct {
	userRepo repository.UserRepo
}

func New(userRepo repository.UserRepo) Service {
	return &service{
		userRepo: userRepo,
	}
}
