package domain

import (
	"time"
)

type UserEntity struct {
	ID        uint      `gorm:"primaryKey"`
	Phone     string    `gorm:"size:20;uniqueIndex;not null"`
	Email     string    `gorm:"size:255;unique;not null"`
	Name      string    `gorm:"size:100"`
	Status    int8      `gorm:"type:tinyint;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
}

func (UserEntity) TableName() string {
	return "users"
}
