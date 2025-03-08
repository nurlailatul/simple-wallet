package domain

import "time"

type BalanceHistoryEntity struct {
	ID              uint      `json:"id" db:"id"`
	WalletID        int64     `json:"wallet_id" db:"wallet_id"`
	TransactionID   int64     `json:"transaction_id" db:"transaction_id"`
	TransactionType int       `json:"transaction_type" db:"transaction_type"`
	OriginAmount    float64   `json:"origin_amount,omitempty" db:"origin_amount"`
	Amount          float64   `json:"amount" db:"amount"`
	OperationType   int       `json:"operation_type" db:"operation_type"`
	FinalAmount     float64   `json:"final_amount,omitempty" db:"final_amount"`
	Notes           string    `json:"notes,omitempty" db:"notes"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}

func (BalanceHistoryEntity) TableName() string {
	return "balance_histories"
}
