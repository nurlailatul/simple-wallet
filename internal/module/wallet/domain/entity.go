package domain

import "time"

type WalletEntity struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null;uniqueIndex"`
	Balance   float64   `gorm:"type:decimal(20,2);default:0.00"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
}

func (WalletEntity) TableName() string {
	return "wallets"
}
