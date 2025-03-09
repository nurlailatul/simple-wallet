package disburse

import (
	"time"

	"github.com/gofrs/uuid"
)

type CreateDisburseRequest struct {
	ReceiverBank          string  `json:"receiver_bank" binding:"required" example:"bca"`
	ReceiverAccountNumber string  `json:"receiver_account_number" binding:"required" example:"123456789"`
	Amount                float64 `json:"amount" binding:"required,gt=0" validate:"min=1000"`
	ReferenceID           string  `json:"reference_id" binding:"required" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
}

type CreateDisburseResponse struct {
	WalletID    uint      `json:"wallet_id"`
	NewBalance  float64   `json:"new_balance"`
	Status      string    `json:"status"`
	ReferenceID uuid.UUID `json:"reference_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type GetWalletBalanceResponse struct {
	WalletID  uint      `json:"wallet_id"`
	Balance   float64   `json:"balance"`
	UpdatedAt time.Time `json:"updated_at"`
}
