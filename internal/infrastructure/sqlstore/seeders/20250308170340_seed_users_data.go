package seeders

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"

	transactionDomain "simple-wallet/internal/module/transaction/domain"
	userDomain "simple-wallet/internal/module/user/domain"
	walletDomain "simple-wallet/internal/module/wallet/domain"
)

func init() {
	seeders = append(seeders, &gormigrate.Migration{
		ID: "20250308170340_seed_users_data",
		Migrate: func(tx *gorm.DB) error {
			users := []userDomain.UserEntity{
				{ID: 1, Phone: "+6281234567890", Name: "Nur Lailatul", Email: "nurlailatul@paper.id", CreatedAt: time.Now()},
				{ID: 2, Phone: "+6281234567891", Name: "Admin Paper ID", Email: "admin@paper.id", CreatedAt: time.Now()},
			}
			tx.Create(&users)

			wallets := []walletDomain.WalletEntity{
				{ID: 1, UserID: 1, Balance: 10000.00, CreatedAt: time.Now()},
				{ID: 2, UserID: 2, Balance: 5000.00, CreatedAt: time.Now()},
			}
			tx.Create(&wallets)

			transactions := []transactionDomain.TransactionEntity{
				{WalletID: 1, Amount: 2000.00, Type: transactionDomain.TypeDebit, Status: transactionDomain.StatusCompleted, ReferenceID: uuid.New().String(), CreatedAt: time.Now()},
				{WalletID: 2, Amount: 1500.00, Type: transactionDomain.TypeCredit, Status: transactionDomain.StatusCompleted, ReferenceID: uuid.New().String(), CreatedAt: time.Now()},
			}
			tx.Create(&transactions)

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			// Empty table
			tx.Exec("TRUNCATE TABLE users")
			tx.Exec("TRUNCATE TABLE wallets")
			tx.Exec("TRUNCATE TABLE transactions")

			return nil
		},
	})
}
