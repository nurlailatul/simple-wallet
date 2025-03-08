package migration

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	migrations = append(migrations, &gormigrate.Migration{
		ID: "20250308133605_create_users",
		Migrate: func(tx *gorm.DB) error {
			type User struct {
				ID        uint      `gorm:"primaryKey"`
				Phone     string    `gorm:"size:20;uniqueIndex;not null"`
				Email     string    `gorm:"size:255;unique;not null"`
				Name      string    `gorm:"size:100"`
				Status    int8      `gorm:"type:tinyint;not null"`
				CreatedAt time.Time `gorm:"autoCreateTime"`
				UpdatedAt time.Time `gorm:"autoCreateTime"`
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
				ID        uint      `gorm:"primaryKey"`
				UserID    uint      `gorm:"not null;uniqueIndex"`
				Balance   float64   `gorm:"type:decimal(20,2);default:0.00"`
				CreatedAt time.Time `gorm:"autoCreateTime"`
				UpdatedAt time.Time `gorm:"autoCreateTime"`
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
			tx.Exec("CREATE INDEX idx_wallets_createdAt ON wallets (created_at)")
			tx.Exec("CREATE INDEX idx_wallets_userId_createdAt ON wallets (user_id, created_at)")
			return nil
		},
	})

	migrations = append(migrations, &gormigrate.Migration{
		ID: "20250308134840_create_transactions",
		Migrate: func(tx *gorm.DB) error {
			type Transaction struct {
				ID          uint      `gorm:"primaryKey"`
				WalletID    uint      `gorm:"not null;index"`
				Amount      float64   `gorm:"type:decimal(20,2);not null"`
				Type        int8      `gorm:"type:tinyint;not null"`
				Status      int8      `gorm:"type:tinyint;not null"`
				ReferenceID string    `gorm:"size:100;unique;not null"`
				CreatedAt   time.Time `gorm:"autoCreateTime"`
				CompletedAt time.Time `gorm:"autoCreateTime"`
				UpdateAt    time.Time `gorm:"autoCreateTime"`
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
			tx.Exec("CREATE INDEX idx_transactions_createdAt ON transactions (created_at)")
			tx.Exec("CREATE INDEX idx_transactions_completedAt ON transactions (completed_at)")
			tx.Exec("CREATE INDEX idx_transactions_walletId_type ON transactions (wallet_id, type)")
			tx.Exec("CREATE INDEX idx_transactions_status_createdAt ON transactions (status, created_at)")
			tx.Exec("CREATE INDEX idx_transactions_status_completedAt ON transactions (status, completed_at)")
			tx.Exec("CREATE INDEX idx_transactions_walletId_type_status ON transactions (wallet_id, type, status)")
			return nil
		},
	})
}
