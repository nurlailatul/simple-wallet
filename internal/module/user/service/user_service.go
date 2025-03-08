package service

import (
	"context"
	"simple-wallet/internal/module/user/domain"
	"simple-wallet/internal/module/user/repository"
)

type UserService struct {
	userRepo repository.UserRepositoryInterface
}

func NewUserService(
	userRepo repository.UserRepositoryInterface,
) UserServiceInterface {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetById(ctx context.Context, userID int64) *domain.UserEntity {
	return s.userRepo.GetById(ctx, userID)
}
