package app

import (
	"context"

	"simple-wallet/config"

	"simple-wallet/internal/infrastructure"
)

type Application struct {
	Infrastructure infrastructure.Infrastructure
}

func NewApp(ctx context.Context) (*Application, error) {
	infrastructure, err := infrastructure.NewInfra(ctx, *config.All())
	if err != nil {
		return nil, err
	}

	app := &Application{
		Infrastructure: infrastructure,
	}

	return app, nil
}
