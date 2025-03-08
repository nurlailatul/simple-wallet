package seeders

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"

	userDomain "simple-wallet/internal/module/user/domain"
	walletDomain "simple-wallet/internal/module/wallet/domain"
)

func init() {
	seeders = append(seeders, &gormigrate.Migration{
		ID: "20250308170340_seed_users_data",
		Migrate: func(tx *gorm.DB) error {
			users := []userDomain.UserEntity{
				{ID: 1, Phone: "+6281234567890", Name: "Nur Lailatul", Email: "nurlailatul@gmail.com", CreatedAt: time.Now()},
				{ID: 2, Phone: "+6281234567891", Name: "Admin Paper ID", Email: "admin@gmail.com", CreatedAt: time.Now()},
			}
			tx.Create(&users)

			wallets := []walletDomain.WalletEntity{
				{ID: 1, UserID: 1, Balance: 10000.00, CreatedAt: time.Now()},
				{ID: 2, UserID: 2, Balance: 5000.00, CreatedAt: time.Now()},
			}
			tx.Create(&wallets)

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			// Empty table
			tx.Exec("TRUNCATE TABLE users")
			tx.Exec("TRUNCATE TABLE wallets")

			return nil
		},
	})
}
