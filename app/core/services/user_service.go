package services

import (
	"context"
	"errors"
	"strings"

	"my-knowledge-sharing/app/core/ports"
)

type UserService struct {
	repo ports.UserRepositoryPort
}

func NewUserService(repo ports.UserRepositoryPort) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, name string) (int64, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return 0, errors.New("name is required")
	}
	if len(name) < 3 {
		return 0, errors.New("name must be at least 3 characters")
	}

	return s.repo.CreateUser(ctx, name)
}

func (s *UserService) GetUser(ctx context.Context, id int64) (string, error) {
	if id <= 0 {
		return "", errors.New("invalid id")
	}

	return s.repo.FindByID(ctx, id)
}