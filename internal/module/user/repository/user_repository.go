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

var selectQuery = `*,  CASE 
		WHEN no_hp_verified = 1 THEN 3
		WHEN status = 10 THEN 2
		ELSE 1
	END AS status`

func (repo *UserRepository) GetById(ctx context.Context, userId int64) *domain.User {
	dbgorm := repo.gorm

	var entity domain.User
	if err := dbgorm.Table("user").WithContext(ctx).Where("id = ?", userId).First(&entity).Error; err != nil {
		return nil
	}

	return &entity
}
