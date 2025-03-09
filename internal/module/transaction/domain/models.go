package domain

import (
	userDomain "simple-wallet/internal/module/user/domain"
)

type DeductBalanceRequest struct {
	UserID                int64
	WalletID              int64
	Amount                float64
	ReceiverBank          string
	ReceiverAccountNumber string
	ReferenceID           string
	User                  *userDomain.UserEntity
}

type DeductBalanceResponse struct {
	WalletID   int64
	NewBalance float64
	Status     Status
	CreatedAt  uint
}
