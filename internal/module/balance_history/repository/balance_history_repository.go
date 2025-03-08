package repository

import (
	"context"

	"simple-wallet/internal/module/balance_history/domain"

	"gorm.io/gorm"
)

type BalanceHistoryRepository struct {
	gorm *gorm.DB
}

func NewBalanceHistoryRepository(
	dbGorm *gorm.DB,
) BalanceHistoryRepositoryInterface {
	return &BalanceHistoryRepository{
		gorm: dbGorm,
	}
}

func (repo *BalanceHistoryRepository) Create(ctx context.Context, entity domain.BalanceHistoryEntity) error {
	dbs := repo.gorm

	return dbs.WithContext(ctx).Create(&entity).Error
}
