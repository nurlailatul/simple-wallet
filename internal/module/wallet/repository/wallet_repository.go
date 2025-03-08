package repository

import (
	"context"

	"simple-wallet/internal/module/wallet/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WalletRepository struct {
	gorm *gorm.DB
}

func NewWalletRepository(
	dbGorm *gorm.DB,
) WalletRepositoryInterface {
	return &WalletRepository{
		gorm: dbGorm,
	}
}

func (repo *WalletRepository) GetByUserID(ctx context.Context, userID int64) *domain.WalletEntity {
	db := repo.gorm

	var entity domain.WalletEntity
	if err := db.Table("wallets").WithContext(ctx).Where("user_id = ?", userID).First(&entity).Error; err != nil {
		return nil
	}

	return &entity
}

func (repo *WalletRepository) GetByUserIDForLocking(ctx context.Context, id int64) *domain.WalletEntity {
	db := repo.gorm

	var entity domain.WalletEntity
	if err := db.Table("wallets").WithContext(ctx).Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", id).First(&entity).Error; err != nil {
		return nil
	}

	return &entity
}

func (repo *WalletRepository) Update(ctx context.Context, entity *domain.WalletEntity) error {
	dbs := repo.gorm

	return dbs.WithContext(ctx).Save(&entity).Error
}
