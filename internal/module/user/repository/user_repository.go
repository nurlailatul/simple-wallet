package repository

import (
	"context"

	"simple-wallet/internal/module/user/domain"
	"simple-wallet/pkg/db"
)

type UserRepository struct {
	gorm *db.GormDBWrapper
}

func NewUserRepository(dbGorm *db.GormDBWrapper) UserRepositoryInterface {
	return &UserRepository{
		gorm: dbGorm,
	}
}

func (repo *UserRepository) GetById(ctx context.Context, userId int64) *domain.UserEntity {
	dbgorm := repo.gorm

	var entity domain.UserEntity
	if err := dbgorm.Table("user").WithContext(ctx).Where("id = ?", userId).First(&entity).Error; err != nil {
		return nil
	}

	return &entity
}
