package service

import (
	"context"
	"simple-wallet/internal/module/wallet/domain"
	"simple-wallet/internal/module/wallet/repository"
)

type WalletService struct {
	repo repository.WalletRepositoryInterface
}

func NewWalletService(
	repo repository.WalletRepositoryInterface,
) WalletServiceInterface {
	return &WalletService{
		repo: repo,
	}
}

func (s *WalletService) GetByUserID(ctx context.Context, userID int64) *domain.WalletEntity {
	return s.repo.GetByUserID(ctx, userID)
}
