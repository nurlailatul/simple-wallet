package repository

import (
	"context"

	"simple-wallet/internal/module/balance_history/domain"
)

//go:generate mockery --name "BalanceHistoryRepositoryInterface" --output "../mocks" --outpkg "mocks"
type BalanceHistoryRepositoryInterface interface {
	Create(ctx context.Context, entity domain.BalanceHistoryEntity) error
}
