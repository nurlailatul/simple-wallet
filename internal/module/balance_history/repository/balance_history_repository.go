package repository

import (
	"context"

	"simple-wallet/internal/module/balance_history/domain"
	"simple-wallet/pkg/db"
)

type BalanceHistoryRepository struct {
	gorm *db.GormDBWrapper
}

func NewBalanceHistoryRepository(
	dbGorm *db.GormDBWrapper,
) BalanceHistoryRepositoryInterface {
	return &BalanceHistoryRepository{
		gorm: dbGorm,
	}
}

func (repo *BalanceHistoryRepository) Create(ctx context.Context, entity domain.BalanceHistoryEntity) error {
	dbs := repo.gorm

	return dbs.WithContext(ctx).Create(&entity).Error
}
