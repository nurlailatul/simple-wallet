package domain

type WalletEntity struct {
	ID        int64   `json:"id" db:"id"`
	UserID    int64   `json:"user_id" db:"user_id"`
	Balance   float64 `json:"balance" db:"balance"`
	CreatedAt uint    `json:"created_at" db:"created_at"`
	UpdatedAt uint    `json:"updated_at" db:"updated_at"`
}

func (WalletEntity) TableName() string {
	return "wallets"
}
