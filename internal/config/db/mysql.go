package db

import (
	"database/sql"
	"gorm.io/gorm"
	"time"

	"gorm.io/driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"pbi/internal/config"
)

type Database struct {
	Raw 	*sql.DB
	Gorm 	*gorm.DB
}

func InitMysql(cfg config.DBConfig) (*Database, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)
	// fmt.Println("DSN:", dsn)

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
    sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
    
    // Konversi int dari config ke time.Duration
    sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)
    sqlDB.SetConnMaxIdleTime(time.Duration(cfg.ConnMaxIdleTime) * time.Second)

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return &Database{
		Gorm: gormDB,
		Raw:  sqlDB,
	}, nil
}