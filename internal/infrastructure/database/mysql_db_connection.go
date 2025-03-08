package database

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"sync"

	"simple-wallet/config"
	"simple-wallet/pkg/db"

	"github.com/jmoiron/sqlx"
)

var dbConnection DBConnection
var once sync.Once

type DBConnection struct {
	DB *db.SqlxDBWrapper
}

func (db *DBConnection) GetConnection() DBConnection {
	config := *config.All()
	once.Do(func() {
		db := getSqlxDBWrapper(config.Database.Main)
		dbConnection = DBConnection{
			DB: db,
		}
	})
	return dbConnection
}

func getSqlxDBWrapper(configDB db.Config) *db.SqlxDBWrapper {
	dbMaster, err := db.DB(configDB, timeout)
	if err != nil {
		log.Println(err)
		return nil
	}
	dbxMaster := sqlx.NewDb(dbMaster, configDB.Driver)
	return db.NewSqlxDBWrapper(dbxMaster, configDB.Driver, timeout)
}

func (db *DBConnection) BeginTx(ctx context.Context, sql *db.SqlxDBWrapper) context.Context {
	tx, _ := sql.BeginTx(ctx, nil)
	ctx = context.WithValue(ctx, Tx, tx)
	ctx = context.WithValue(ctx, Db, sql)
	return ctx
}

func (db *DBConnection) CommitTx(ctx context.Context) error {
	tx, ok := ctx.Value(Tx).(*sql.Tx)
	if !ok {
		return errors.New("failed to commit on non transaction mode")
	}

	return tx.Commit()
}

func (db *DBConnection) RollbackTx(ctx context.Context) error {
	tx, ok := ctx.Value(Tx).(*sql.Tx)
	if !ok {
		return errors.New("failed to rollback on non transaction mode")
	}
	_ = tx.Rollback()
	return nil
}

func (db *DBConnection) Ping(ctx context.Context, sql *db.SqlxDBWrapper) error {
	err := sql.PingContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (db *DBConnection) RollbackIfError(ctx context.Context, err error) error {
	if err == nil {
		return nil
	}

	if p := recover(); p != nil {
		db.RollbackTx(ctx)
		panic(p)
	} else {
		err = db.RollbackTx(ctx)

		if err == nil {
			return nil
		}
	}

	return errors.New("rollback is failed")
}

func (db *DBConnection) CommitIfNoError(ctx context.Context, err error) error {
	if err != nil {
		return err
	}

	err = db.CommitTx(ctx)
	if err == nil {
		return nil
	}

	return errors.New("commit is failed")

}

func (db *DBConnection) CloseTx(ctx context.Context, err error) {
	if p := recover(); p != nil {
		db.RollbackTx(ctx)
		panic(p)
	} else if err != nil {
		db.RollbackTx(ctx)
	} else {
		db.CommitTx(ctx)
	}
}
