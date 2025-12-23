package container

import (
	"log"

	"pbi/internal/config"
	"pbi/internal/config/db"
	// "pbi/internal/pkg/repository"
	// "pbi/internal/pkg/usecase"
	"pbi/internal/utils"
)

type Container struct {
	Config   *config.Config
	DB       *db.Database
}


func InitContainer() *Container {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("failed load config:", err)
	}

	utils.InitJWT(
		cfg.JWT.Secret,
		cfg.JWT.ExpireDur,
	)

	database, err := db.InitMysql(cfg.DB)
	if err != nil {
		log.Fatal("failed connect db:", err)
	}

	// repositories


	// usecases


	return &Container{
		Config: cfg,
		DB:     database,
	}
}