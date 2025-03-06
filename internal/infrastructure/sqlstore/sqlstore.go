package sqlstore

import (
	"context"
	"fmt"
	"time"

	"simple-wallet/pkg/db"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	goLogger "gorm.io/gorm/logger"
)

type Store interface {
	SqlxDB() *db.SqlxDBWrapper
	GormDb() *db.GormDBWrapper
	Ping(ctx context.Context) error
}

type SQLStore struct {
	sqlxDB            *db.SqlxDBWrapper
	gormDB            *db.GormDBWrapper
	driver            string
	connectionTimeout time.Duration
}

func NewSQLStore(ctx context.Context, dbConfig db.Config) (Store, error) {
	connectionTimeout := 3 * time.Second

	log.Info(ctx,
		fmt.Sprintf("database configs master, max open %d, max idle %d, max lifetime %d",
			dbConfig.MaxOpen, dbConfig.MaxIdle, dbConfig.MaxLifetime))

	configMaster := db.Config{
		Host:        dbConfig.Host,
		Port:        dbConfig.Port,
		User:        dbConfig.User,
		Password:    dbConfig.Password,
		Name:        dbConfig.Name,
		MaxOpen:     int(dbConfig.MaxOpen),
		MaxIdle:     int(dbConfig.MaxIdle),
		MaxLifetime: int(dbConfig.MaxLifetime),
		MaxIdleTime: int(dbConfig.MaxIdleTime),
		// CA:          dbCa,
		ServerName: dbConfig.ServerName,
		Location:   dbConfig.Location,
		// ParseTime:  true,
		Driver: dbConfig.Driver,
	}

	dbMaster, err := db.DB(configMaster, connectionTimeout)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	dbxMaster := sqlx.NewDb(dbMaster, "mysql")

	dbGormMaster, err := gorm.Open(mysql.Open(db.DataSourceName(configMaster)), &gorm.Config{
		PrepareStmt: true,
		Logger:      goLogger.Default.LogMode(goLogger.Silent),
	})
	if err != nil {
		log.Println(err)
		log.Printf("Config master: %+v\n", configMaster)
		return nil, err
	}

	masterWrapper := db.NewSqlxDBWrapper(dbxMaster, dbConfig.Driver, connectionTimeout)

	masterGormWrapper := db.NewGormDBWrapper(dbGormMaster, dbConfig.Driver, connectionTimeout)

	return &SQLStore{
		sqlxDB:            masterWrapper,
		gormDB:            masterGormWrapper,
		driver:            dbConfig.Driver,
		connectionTimeout: connectionTimeout,
	}, nil
}

func (s *SQLStore) SqlxDB() *db.SqlxDBWrapper {
	return s.sqlxDB
}

func (s *SQLStore) GormDb() *db.GormDBWrapper {
	return s.gormDB
}

func (s *SQLStore) Ping(ctx context.Context) error {
	err := s.sqlxDB.PingContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
