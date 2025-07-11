package sql

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseOptions struct {
	DriverName      string
	Host            string
	Port            int
	Username        string
	Password        string
	Database        string
	SSLMode         string
	Timeout         int
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int
}

func NewDatabaseOptions() *DatabaseOptions {
	return &DatabaseOptions{
		DriverName:      "postgres",
		Host:            "localhost",
		Port:            5432,
		Username:        "postgres",
		Password:        "password",
		Database:        "mydb",
		SSLMode:         "disable",
		Timeout:         30,
		MaxOpenConns:    100,
		MaxIdleConns:    10,
		ConnMaxLifetime: 300,
	}
}

func (opts *DatabaseOptions) GetDBInstance() (*gorm.DB, error) {
	if opts.DriverName == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			opts.Username, opts.Password, opts.Host, opts.Port, opts.Database)

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}

		sqlDB, err := db.DB()
		if err != nil {
			return nil, err
		}
		sqlDB.SetMaxOpenConns(opts.MaxOpenConns)
		sqlDB.SetMaxIdleConns(opts.MaxIdleConns)
		sqlDB.SetConnMaxLifetime(time.Duration(opts.ConnMaxLifetime) * time.Second)

		return db, nil
	}

	if opts.DriverName == "postgres" {
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s connect_timeout=%d",
			opts.Host, opts.Port, opts.Username, opts.Password, opts.Database, opts.SSLMode, opts.Timeout)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}

		return db, nil
	}

	return nil, fmt.Errorf("unsupported driver: %s", opts.DriverName)
}
