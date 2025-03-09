package domain

import (
	userDomain "simple-wallet/internal/module/user/domain"
)

type DeductBalanceRequest struct {
	UserID                int64   `json:"user_id"`
	WalletID              int64   `json:"wallet_id"`
	Amount                float64 `json:"amount"`
	ReceiverBank          string  `json:"receiver_bank"`
	ReceiverAccountNumber string  `json:"receiver_account_number"`
	ReferenceID           string  `json:"reference_id"`
	User                  *userDomain.UserEntity
}
