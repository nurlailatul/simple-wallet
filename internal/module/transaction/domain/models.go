package domain

type DeductBalanceRequest struct {
	WalletID              int64   `json:"wallet_id"`
	Amount                float64 `json:"amount"`
	ReceiverBank          string  `json:"receiver_bank"`
	ReceiverAccountNumber string  `json:"receiver_account_number"`
	ReferenceID           string  `json:"reference_id"`
}
