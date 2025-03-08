package repository

import (
	"context"

	"simple-wallet/internal/module/user/domain"
)

//go:generate mockery --name "UserRepositoryInterface" --output "../mocks" --outpkg "mocks"
type UserRepositoryInterface interface {
	GetById(ctx context.Context, userID int64) *domain.UserEntity
}
