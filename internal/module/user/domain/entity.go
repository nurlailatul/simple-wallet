package domain

type User struct {
	ID       int64  `gorm:"primaryKey" json:"id"`
	Username string `gorm:"not null; unique" json:"username"`
}

func (User) TableName() string {
	return "user"
}
