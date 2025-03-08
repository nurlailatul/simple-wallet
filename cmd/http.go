package cmd

import (
	"context"
	"fmt"
	"runtime"

	"simple-wallet/config"
	"simple-wallet/internal/app"
	"simple-wallet/internal/app/server"

	log "github.com/sirupsen/logrus"
	"gopkg.in/ukautz/clif.v1"
)

func (c *Console) StartServer() *clif.Command {
	return clif.NewCommand("start", "Starting HTTP server", func(o *clif.Command, in clif.Input, out clif.Output) error {

		ctx, cancel := context.WithCancel(context.Background())
		defer func() {
			go callOnInterrupt(cancel)
		}()

		log.Info(ctx, "Runtime Go version "+runtime.Version())

		err := serveHTTP(ctx)
		if err != nil {
			return err
		}

		return nil
	})
}

func serveHTTP(ctx context.Context) error {
	cfg := config.All()

	httpServer := server.New(
		server.WithPort(fmt.Sprintf(":%d", cfg.Server.Port)),
	)

	httpApp, err := app.NewApp(ctx)
	if err != nil {
		return err
	}

	appServices := httpApp.SetupDependencies(ctx)
	// middleware := appServices.SetupMiddleware()
	route := appServices.SetupHttpRouteHandler(cfg)
	httpServer.RegisterRoutes(route)

	if err := httpServer.Start(); err != nil {
		log.Fatal(ctx, "cannot start server with error", err)
	}

	return nil
}
