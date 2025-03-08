package domain

import (
	"time"
)

type UserEntity struct {
	ID        uint      `json:"id" db:"id"`
	Phone     string    `json:"phone" db:"phone"`
	Email     string    `json:"email" db:"email"`
	Name      string    `json:"name" db:"name"`
	Status    int8      `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (UserEntity) TableName() string {
	return "users"
}
