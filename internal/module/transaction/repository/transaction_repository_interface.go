package repository

import (
	"context"

	"simple-wallet/internal/module/transaction/domain"
)

//go:generate mockery --name "TransactionRepositoryInterface" --output "../mocks" --outpkg "mocks"
type TransactionRepositoryInterface interface {
	GetByID(ctx context.Context, id int64) *domain.TransactionEntity
	GetByReferenceID(ctx context.Context, referenceID string) *domain.TransactionEntity
	Create(ctx context.Context, entity domain.TransactionEntity) (id int64, err error)
}
