package app

import (
	"context"

	"simple-wallet/config"
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
	SwaggerAuthMiddleware *middleware.AuthMiddleware
	ApiV1AuthMiddleware   *middleware.AuthMiddleware

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

	swaggerAuthMiddleware := middleware.NewAuthMiddleware(cfg.Swagger.ApiKey)
	apiV1AuthMiddleware := middleware.NewAuthMiddleware(cfg.App.ApiKey)

	userRepo := userRepository.NewUserRepository(gormWrapper)
	transactionRepo := transactionRepository.NewTransactionRepository(gormWrapper.DB)
	walletRepo := walletRepository.NewWalletRepository(gormWrapper.DB)

	userService := userService.NewUserService(userRepo)
	walletService := walletService.NewWalletService(walletRepo)
	transactionService := transactionService.NewTransactionService(
		transactionRepo,
		gormWrapper,
	)

	return &AppService{
		SwaggerAuthMiddleware: swaggerAuthMiddleware,
		ApiV1AuthMiddleware:   apiV1AuthMiddleware,

		UserRepo: userRepo,

		TransactionService: transactionService,
		UserService:        userService,
		WalletService:      walletService,
	}
}
