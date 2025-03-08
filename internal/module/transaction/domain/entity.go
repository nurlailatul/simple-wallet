package domain

import "time"

type Status int8

const (
	StatusPending   Status = 0
	StatusCompleted Status = 1
	StatusFailed    Status = 2
)

type Type int8

const (
	TypeDebit  Type = 1
	TypeCredit Type = 2
)

type TransactionEntity struct {
	ID          uint      `gorm:"primaryKey"`
	WalletID    uint      `gorm:"not null;index"`
	Amount      float64   `gorm:"type:decimal(20,2);not null"`
	Type        Type      `gorm:"type:tinyint;not null"`
	Status      Status    `gorm:"type:tinyint;not null"`
	ReferenceID string    `gorm:"size:100;unique;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	CompletedAt time.Time `gorm:"autoCreateTime"`
	UpdateAt    time.Time `gorm:"autoCreateTime"`
}

func (TransactionEntity) TableName() string {
	return "transactions"
}
