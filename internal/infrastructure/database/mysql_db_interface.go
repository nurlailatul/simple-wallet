package database

import (
	"context"

	"simple-wallet/pkg/db"
)

type MySQLDBInterface interface {
	GetConnection() DBConnection
	BeginTx(ctx context.Context, sql *db.SqlxDBWrapper) context.Context
	CommitTx(ctx context.Context) error
	RollbackTx(ctx context.Context) error
	Ping(ctx context.Context, sql *db.SqlxDBWrapper) error
	RollbackIfError(ctx context.Context, err error) error
	CommitIfNoError(ctx context.Context, err error) error
	CloseTx(ctx context.Context, err error)
}
