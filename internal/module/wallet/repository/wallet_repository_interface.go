package repository

import (
	"context"

	"simple-wallet/internal/module/wallet/domain"
)

//go:generate mockery --name "WalletRepositoryInterface" --output "../mocks" --outpkg "mocks"
type WalletRepositoryInterface interface {
	GetByUserID(ctx context.Context, userID int64) *domain.WalletEntity
	GetByUserIDForLocking(ctx context.Context, userID int64) *domain.WalletEntity
	Update(ctx context.Context, entity *domain.WalletEntity) error
}
