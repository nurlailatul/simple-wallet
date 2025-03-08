package repository

import (
	"context"

	"simple-wallet/internal/module/wallet/domain"
	"simple-wallet/pkg/db"
)

type WalletRepository struct {
	gorm *db.GormDBWrapper
	sqlx *db.SqlxDBWrapper
}

func NewWalletRepository(
	dbGorm *db.GormDBWrapper,
	dbSqlx *db.SqlxDBWrapper,
) WalletRepositoryInterface {
	return &WalletRepository{
		gorm: dbGorm,
		sqlx: dbSqlx,
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
	db := repo.sqlx

	var entity domain.WalletEntity
	query := "SELECT * FROM wallets WHERE id = ? FOR UPDATE"
	if err := db.GetContext(ctx, &entity, query, id); err != nil {
		return nil
	}

	return &entity
}

func (repo *WalletRepository) Update(ctx context.Context, entity *domain.WalletEntity) error {
	dbs := repo.gorm

	return dbs.WithContext(ctx).Save(&entity).Error
}
