package main

import (
	"gin/internal/infra/db"
)

func main() {
	db.InitMysql()
}