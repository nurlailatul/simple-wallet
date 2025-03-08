package repository

import (
	"context"

	"simple-wallet/internal/module/transaction/domain"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	gorm *gorm.DB
}

func NewTransactionRepository(dbGorm *gorm.DB) TransactionRepositoryInterface {
	return &TransactionRepository{
		gorm: dbGorm,
	}
}

func (repo *TransactionRepository) GetByID(ctx context.Context, id int64) *domain.TransactionEntity {
	dbgorm := repo.gorm

	var entity domain.TransactionEntity
	if err := dbgorm.Table("transactions").WithContext(ctx).Where("id = ?", id).First(&entity).Error; err != nil {
		return nil
	}

	return &entity
}

func (repo *TransactionRepository) GetByReferenceID(ctx context.Context, referenceID string) *domain.TransactionEntity {
	dbgorm := repo.gorm

	var entity domain.TransactionEntity
	if err := dbgorm.Table("transactions").WithContext(ctx).Where("reference_id = ?", referenceID).First(&entity).Error; err != nil {
		return nil
	}

	return &entity
}

func (repo *TransactionRepository) Create(ctx context.Context, entity domain.TransactionEntity) (id int64, err error) {
	dbs := repo.gorm

	if err := dbs.WithContext(ctx).Create(&entity).Error; err != nil {
		return 0, err
	}

	return entity.ID, nil
}
