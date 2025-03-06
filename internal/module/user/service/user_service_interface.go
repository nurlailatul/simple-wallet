package service

import (
	"context"
	"simple-wallet/internal/module/user/domain"
)

//go:generate mockery --name "UserServiceInterface" --output "../mocks" --outpkg "mocks"
type UserServiceInterface interface {
	GetById(ctx context.Context, userID int64) *domain.User
}
