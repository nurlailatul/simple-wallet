package domain

import "time"

type Status int8

const (
	StatusPending   Status = 0
	StatusCompleted Status = 1
	StatusFailed    Status = 2
)

type TransactionEntity struct {
	ID                    int64      `json:"id" db:"id"`
	WalletID              int64      `json:"wallet_id" db:"wallet_id"`
	Amount                float64    `json:"amount" db:"amount"`
	ReceiverBank          string     `json:"receiver_bank" db:"receiver_bank"`
	ReceiverAccountNumber string     `json:"receiver_account_number" db:"receiver_account_number"`
	Status                Status     `json:"status" db:"status"`
	ReferenceID           string     `json:"reference_id" db:"reference_id"`
	CreatedAt             time.Time  `json:"created_at" db:"created_at"`
	CompletedAt           *time.Time `json:"completed_at" db:"completed_at"`
	UpdatedAt             time.Time  `json:"updated_at" db:"updated_at"`
}

func (TransactionEntity) TableName() string {
	return "transactions"
}
