package infrastructure

import (
	"context"

	"simple-wallet/config"
	"simple-wallet/internal/infrastructure/sqlstore"
)

// Infrastructure is the wrapper for infra dependencies.
type Infrastructure interface {
	DB() sqlstore.Store
}

type Infra struct {
	db sqlstore.Store
}

func NewInfra(ctx context.Context, config config.Configuration) (Infrastructure, error) {
	// init sql store
	db, err := sqlstore.NewSQLStore(ctx, config.Database.Main)
	if err != nil {
		return nil, err
	}

	// init others
	return &Infra{
		db: db,
	}, nil
}

func (i *Infra) DB() sqlstore.Store {
	return i.db
}
