package app

import (
	"context"

	"simple-wallet/config"
	"simple-wallet/internal/middleware"
	userRepository "simple-wallet/internal/module/user/repository"
	userService "simple-wallet/internal/module/user/service"
)

type AppService struct {
	SwaggerAuthMiddleware *middleware.SwaggerAuthMiddleware

	UserRepo    userRepository.UserRepositoryInterface
	UserService userService.UserServiceInterface
}

func (a *Application) SetupDependencies(ctx context.Context, cfg *config.Configuration) *AppService {
	masterGormWrapper := a.Infrastructure.DB().GormDb()

	swaggerAuthMiddleware := middleware.NewSwaggerAuthMiddleware(&cfg.Swagger)

	userRepo := userRepository.NewUserRepository(masterGormWrapper)

	userService := userService.NewUserService(userRepo)

	return &AppService{
		SwaggerAuthMiddleware: swaggerAuthMiddleware,

		UserRepo:    userRepo,
		UserService: userService,
	}
}
