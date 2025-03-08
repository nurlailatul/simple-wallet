package service

import (
	"context"
	"simple-wallet/internal/module/wallet/domain"
)

//go:generate mockery --name "WalletServiceInterface" --output "../mocks" --outpkg "mocks"
type WalletServiceInterface interface {
	GetByUserID(ctx context.Context, userID int64) *domain.WalletEntity
}
