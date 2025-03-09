package wallets

import (
	transactionDomain "simple-wallet/internal/module/transaction/domain"
	"time"
)

type CreateDisburseRequest struct {
	ReceiverBank          string  `json:"receiver_bank" binding:"required" example:"bca"`
	ReceiverAccountNumber string  `json:"receiver_account_number" binding:"required" example:"123456789"`
	Amount                float64 `json:"amount" binding:"required,gt=0" validate:"min=1000"`
	ReferenceID           string  `json:"reference_id" binding:"required" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
}

type CreateDisburseResponse struct {
	WalletID              int64   `json:"wallet_id"`
	ReceiverBank          string  `json:"receiver_bank" example:"bca"`
	ReceiverAccountNumber string  `json:"receiver_account_number" example:"123456789"`
	Amount                float64 `json:"amount" validate:"min=1000"`
	NewBalance            float64 `json:"new_balance"`
	Status                int8    `json:"status"`
	ReferenceID           string  `json:"reference_id"`
	CreatedAt             string  `json:"created_at"`
}

func responseMapping(entity transactionDomain.DeductBalanceResponse, req CreateDisburseRequest) CreateDisburseResponse {
	unixTime := int64(entity.CreatedAt)
	t := time.Unix(unixTime, 0)
	formattedDate := t.Format("2006-01-02 15:04:05")
	return CreateDisburseResponse{
		WalletID:              entity.WalletID,
		ReceiverBank:          req.ReceiverBank,
		ReceiverAccountNumber: req.ReceiverAccountNumber,
		Amount:                req.Amount,
		NewBalance:            entity.NewBalance,
		Status:                int8(entity.Status),
		ReferenceID:           req.ReferenceID,
		CreatedAt:             formattedDate,
	}
}
