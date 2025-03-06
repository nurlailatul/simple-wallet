package db

import (
	"database/sql"
	"fmt"
	"net/url"
	"time"
)

const (
	DefaultMaxOpen     = 50
	DefaultMaxIdle     = 10
	DefaultMaxLifetime = 3
)

type Config struct {
	Driver      string `mapstructure:"driver"`
	Host        string `mapstructure:"host"`
	Port        string `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Name        string `mapstructure:"name"`
	MaxOpen     int    `mapstructure:"maxopen"`
	MaxIdle     int    `mapstructure:"maxidle"`
	MaxLifetime int    `mapstructure:"maxlifetime"` // in minutes
	MaxIdleTime int    // in minutes
	CA          []byte
	ServerName  string
	ParseTime   bool
	Location    string
}

func DataSourceName(config Config) string {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.Name)
	val := url.Values{}

	if config.ParseTime {
		val.Add("parseTime", "1")
	}
	if len(config.Location) > 0 {
		val.Add("loc", config.Location)
	}

	if len(val) == 0 {
		return connection
	}
	return fmt.Sprintf("%s?%s", connection, val.Encode())
}

// DB return new sql db
func DB(config Config, timeout time.Duration) (*sql.DB, error) {

	db, err := sql.Open(config.Driver, DataSourceName(config))
	if err != nil {
		return nil, err
	}

	if config.MaxOpen > 0 {
		db.SetMaxOpenConns(config.MaxOpen)
	} else {
		db.SetMaxOpenConns(DefaultMaxOpen)
	}

	if config.MaxIdle > 0 {
		db.SetMaxIdleConns(config.MaxIdle)
	} else {
		db.SetMaxIdleConns(DefaultMaxIdle)
	}

	if config.MaxLifetime > 0 {
		db.SetConnMaxLifetime(time.Duration(config.MaxLifetime) * time.Minute)
	} else {
		db.SetConnMaxLifetime(time.Duration(DefaultMaxLifetime) * time.Minute)
	}

	if config.MaxIdleTime > 0 {
		db.SetConnMaxIdleTime(time.Duration(config.MaxIdleTime) * time.Minute)
	}

	return db, nil
}
