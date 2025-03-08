package service

import (
	"context"
	"simple-wallet/internal/module/transaction/domain"
)

//go:generate mockery --name "TransactionServiceInterface" --output "../mocks" --outpkg "mocks"
type TransactionServiceInterface interface {
	GetByReferenceID(ctx context.Context, referenceID string) *domain.TransactionEntity
	DeductBalance(ctx context.Context, request domain.DeductBalanceRequest) error
}
