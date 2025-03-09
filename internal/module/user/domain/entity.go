package domain

type UserEntity struct {
	ID        int64  `json:"id" db:"id"`
	Phone     string `json:"phone" db:"phone"`
	Email     string `json:"email" db:"email"`
	Name      string `json:"name" db:"name"`
	Status    int8   `json:"status" db:"status"`
	CreatedAt uint   `json:"created_at" db:"created_at"`
	UpdatedAt uint   `json:"updated_at" db:"updated_at"`
}

func (UserEntity) TableName() string {
	return "users"
}
