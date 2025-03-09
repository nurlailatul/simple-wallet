package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	migrations = append(migrations, &gormigrate.Migration{
		ID: "20250308133605_create_users",
		Migrate: func(tx *gorm.DB) error {
			type User struct {
				ID        uint   `gorm:"primaryKey"`
				Phone     string `gorm:"size:20;uniqueIndex;not null"`
				Email     string `gorm:"size:255;unique;not null"`
				Name      string `gorm:"size:100"`
				Status    int8   `gorm:"type:tinyint;not null"`
				CreatedAt uint   `gorm:"autoCreateTime"`
				UpdatedAt uint   `gorm:"autoCreateTime"`
			}

			return tx.AutoMigrate(&User{})
		},

		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("users")
		},
	})

	migrations = append(migrations, &gormigrate.Migration{
		ID: "20250308133615_create_index_users",
		Migrate: func(tx *gorm.DB) error {
			tx.Exec("CREATE INDEX idx_users_phone_email ON users (phone, email)")
			return nil
		},
	})

	migrations = append(migrations, &gormigrate.Migration{
		ID: "20250308134529_create_wallets",
		Migrate: func(tx *gorm.DB) error {
			type Wallet struct {
				ID        uint    `gorm:"primaryKey"`
				UserID    uint    `gorm:"not null;uniqueIndex"`
				Balance   float64 `gorm:"type:decimal(20,2);default:0.00"`
				CreatedAt uint    `gorm:"autoCreateTime;index"`
				UpdatedAt uint    `gorm:"autoCreateTime"`
			}

			return tx.AutoMigrate(&Wallet{})
		},

		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("wallets")
		},
	})

	migrations = append(migrations, &gormigrate.Migration{
		ID: "20250308134539_create_index_wallets",
		Migrate: func(tx *gorm.DB) error {
			tx.Exec("CREATE INDEX idx_wallets_userId_createdAt ON wallets (user_id, created_at)")
			return nil
		},
	})

	migrations = append(migrations, &gormigrate.Migration{
		ID: "20250308134840_create_transactions",
		Migrate: func(tx *gorm.DB) error {
			type Transaction struct {
				ID                    uint    `gorm:"primaryKey"`
				WalletID              uint    `gorm:"not null;index"`
				Amount                float64 `gorm:"type:decimal(20,2);not null"`
				ReceiverBank          string  `gorm:"size:100;not null"`
				ReceiverAccountNumber string  `gorm:"size:100;not null"`
				Status                int8    `gorm:"type:tinyint;not null"`
				ReferenceID           string  `gorm:"size:100;unique;not null"`
				CreatedAt             uint    `gorm:"autoCreateTime;index"`
				CompletedAt           *uint   `gorm:"autoCreateTime;index"`
				UpdatedAt             uint    `gorm:"autoCreateTime"`
			}

			return tx.AutoMigrate(&Transaction{})
		},

		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("transactions")
		},
	})

	migrations = append(migrations, &gormigrate.Migration{
		ID: "20250308140557_create_index_transactions",
		Migrate: func(tx *gorm.DB) error {
			tx.Exec("CREATE INDEX idx_transactions_status_createdAt ON transactions (status, created_at)")
			tx.Exec("CREATE INDEX idx_transactions_status_completedAt ON transactions (status, completed_at)")
			tx.Exec("CREATE INDEX idx_transactions_walletId_status ON transactions (wallet_id, status)")
			tx.Exec("CREATE INDEX idx_transactions_receiver_bank_receiver_account_number ON transactions (receiver_bank, receiver_account_number)")
			return nil
		},
	})

	migrations = append(migrations, &gormigrate.Migration{
		ID: "20250308203422_create_balance_histories",
		Migrate: func(tx *gorm.DB) error {
			type BalanceHistory struct {
				ID              uint    `gorm:"primaryKey;autoIncrement"`
				WalletID        uint    `gorm:"index"`
				TransactionID   uint    `gorm:"not null"`
				TransactionType int     `gorm:"not null"`
				OriginAmount    float64 `gorm:"default:null"`
				Amount          float64 `gorm:"not null"`
				OperationType   int     `gorm:"not null"`
				FinalAmount     float64 `gorm:"default:null"`
				Notes           string  `gorm:"type:text"`
				CreatedAt       uint    `gorm:"autoCreateTime"`
			}

			return tx.AutoMigrate(&BalanceHistory{})
		},

		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("balance_histories")
		},
	})

	migrations = append(migrations, &gormigrate.Migration{
		ID: "20250308203432_create_index_balance_histories",
		Migrate: func(tx *gorm.DB) error {
			tx.Exec("CREATE INDEX idx_balance_histories_wallet_id_created_at ON transactions (wallet_id, created_at)")
			return nil
		},
	})
}
