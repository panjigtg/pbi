package db

import (
	"database/sql"
	"os"
	"time"
	"errors"	

	"github.com/joho/godotenv"
)

func InitMysql() (*sql.DB, error) {
	// 1. Load file .env
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// 2. Ambil nilai dari variabel
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbName == "" {
		return nil, errors.New("db env missing")
	}

	// 3. Bangun DSN (Data Source Name)
	dsn := dbUser + ":" +
		dbPassword + "@tcp(" +
		dbHost + ":" +
		dbPort + ")/" +
		dbName +
		"?parseTime=true&charset=utf8mb4&loc=Local"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 5)

	return db, nil
}