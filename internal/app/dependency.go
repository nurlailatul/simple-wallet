package app

import (
	"context"

	"simple-wallet/config"
	"simple-wallet/internal/infrastructure/database"
	"simple-wallet/internal/middleware"
	balanceHistoryRepository "simple-wallet/internal/module/balance_history/repository"
	transactionRepository "simple-wallet/internal/module/transaction/repository"
	transactionService "simple-wallet/internal/module/transaction/service"
	userRepository "simple-wallet/internal/module/user/repository"
	userService "simple-wallet/internal/module/user/service"
	walletRepository "simple-wallet/internal/module/wallet/repository"
	walletService "simple-wallet/internal/module/wallet/service"
)

type AppService struct {
	SwaggerAuthMiddleware *middleware.SwaggerAuthMiddleware

	UserRepo           userRepository.UserRepositoryInterface
	UserService        userService.UserServiceInterface
	TransactionService transactionService.TransactionServiceInterface
	TransactionRepo    transactionRepository.TransactionRepositoryInterface
	BalanceHistoryRepo balanceHistoryRepository.BalanceHistoryRepositoryInterface
	WalletRepo         walletRepository.WalletRepositoryInterface
	WalletService      walletService.WalletServiceInterface
}

func (a *Application) SetupDependencies(ctx context.Context, cfg *config.Configuration) *AppService {
	gormWrapper := a.Infrastructure.DB().GormDb()
	sqlxWrapper := a.Infrastructure.DB().SqlxDB()
	connection := database.DBConnection{}
	connection = connection.GetConnection()

	swaggerAuthMiddleware := middleware.NewSwaggerAuthMiddleware(&cfg.Swagger)

	userRepo := userRepository.NewUserRepository(gormWrapper)
	balanceHistoryRepo := balanceHistoryRepository.NewBalanceHistoryRepository(gormWrapper)
	transactionRepo := transactionRepository.NewTransactionRepository(gormWrapper)
	walletRepo := walletRepository.NewWalletRepository(gormWrapper, sqlxWrapper)

	userService := userService.NewUserService(userRepo)
	walletService := walletService.NewWalletService(walletRepo)
	transactionService := transactionService.NewTransactionService(
		balanceHistoryRepo,
		transactionRepo,
		walletRepo,
		&connection,
	)

	return &AppService{
		SwaggerAuthMiddleware: swaggerAuthMiddleware,

		UserRepo: userRepo,

		TransactionService: transactionService,
		UserService:        userService,
		WalletService:      walletService,
	}
}
