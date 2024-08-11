package service

import (
	"context"
	"errors"
	"time"

	"github.com/chanudinho/go-api-example/pkg/user/domain"
	"github.com/chanudinho/go-api-example/pkg/user/models"
	"github.com/golang-jwt/jwt/v5"
)

func (s *service) Login(ctx context.Context, user domain.User) (string, error) {
	u, err := s.userRepo.FindByEmail(ctx, user.Email)
	if err != nil {
		return "", err
	}

	if !u.ComparePassword(user.Password) {
		return "", errors.New("invalid password")
	}

	token, err := s.createToken(u)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *service) createToken(user models.User) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"role":   user.Role,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}).SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return token, nil
}
