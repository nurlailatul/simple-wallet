package app

import (
	"context"

	userRepository "simple-wallet/internal/module/user/repository"
	userService "simple-wallet/internal/module/user/service"
)

type AppService struct {
	UserRepo    userRepository.UserRepositoryInterface
	UserService userService.UserServiceInterface
}

func (a *Application) SetupDependencies(ctx context.Context) *AppService {

	masterGormWrapper := a.Infrastructure.DB().GormDb()

	userRepo := userRepository.NewUserRepository(masterGormWrapper)

	userService := userService.NewUserService(userRepo)

	return &AppService{
		UserRepo:    userRepo,
		UserService: userService,
	}
}
