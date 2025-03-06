package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"simple-wallet/config"
	"simple-wallet/internal/app/server"

	log "github.com/sirupsen/logrus"
)

func StartServer() {

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		go callOnInterrupt(cancel)
	}()

	log.Info(ctx, "Runtime Go version "+runtime.Version())

	err := serveHTTP(ctx)
	if err != nil {
		log.Info(ctx, "Error serve HTTP  ", err)
	}

}

func serveHTTP(ctx context.Context) error {
	cfg := config.All()

	httpServer := server.New(
		server.WithPort(fmt.Sprintf(":%d", cfg.Server.Port)),
		// server.WithMiddleware(server.CORS()),
	)

	app, err := NewApp(ctx)
	if err != nil {
		return err
	}

	appServices := app.SetupDependencies(ctx)
	route := appServices.SetupHttpRouteHandler(cfg)
	httpServer.RegisterRoutes(route)

	if err := httpServer.Start(); err != nil {
		log.Fatal(ctx, "cannot start server with error", err)
	}

	return nil
}

func callOnInterrupt(cancel context.CancelFunc) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	<-sigCh
	cancel()
}
