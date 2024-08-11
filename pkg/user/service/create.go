package service

import (
	"context"

	"github.com/chanudinho/go-api-example/pkg/user/domain"
)

func (s *service) Create(ctx context.Context, userToCreate domain.User) (domain.User, error) {
	user, err := s.userRepo.Create(ctx, userToCreate.ToModel())
	if err != nil {
		return domain.User{}, err
	}

	var userDomain domain.User
	userDomain.FromModel(user)
	return userDomain, nil
}
