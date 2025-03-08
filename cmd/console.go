package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"simple-wallet/internal/infrastructure/sqlstore"
)

type Console struct {
	Migrator *sqlstore.Migrator
}

func Init() (*Console, error) {
	sqlMigrator, err := sqlstore.NewMigrator()
	if err != nil {
		return nil, err
	}

	return &Console{
		Migrator: sqlMigrator,
	}, nil
}

func callOnInterrupt(cancel context.CancelFunc) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	<-sigCh
	cancel()
}
