package container

import (
	"log"

	"pbi/internal/config"
	"pbi/internal/config/db"
	"pbi/internal/pkg/repository"
	"pbi/internal/pkg/usecase"
	"pbi/internal/utils"
)

type Container struct {
	Config  *config.Config
	DB      *db.Database
	AuthUsc usecase.AuthUseCase
	UserUsc usecase.UserUseCase
	CUsc 	usecase.CategoryUseCase
	AddrUsc usecase.AddressUsecase
	TokoUsc usecase.TokoUsecase
	DestUsc usecase.DestinationUsecase
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
	userRepo 			:= repository.NewUserRepository(database.Raw)
	tokoRepo			:= repository.NewTokoRepository(database.Raw)
	categoryRepo 		:= repository.NewCategoryRepository(database.Raw)
	addressRepo 		:= repository.NewAddressRepository()
	destinationRepo 	:= repository.NewDestinationRepo(database.Gorm)

	// usecases
	addressUsc 			:= usecase.NewAddressUsecase(addressRepo)
	authUsc 			:= usecase.NewAuthUseCase(userRepo, tokoRepo, addressUsc)
	userUsc 			:= usecase.NewUserUseCase(userRepo, addressUsc)
	categoryUsc 		:= usecase.NewCategoryUseCase(categoryRepo)
	tokoUsc 			:= usecase.NewTokoUsecase(tokoRepo)
	destUsc				:= usecase.NewDestinationUsecase(destinationRepo)


	return &Container{
		Config:  cfg,
		DB:      database,
		AuthUsc: authUsc,
		UserUsc: userUsc,
		CUsc:    categoryUsc,
		AddrUsc: addressUsc,
		TokoUsc: tokoUsc,
		DestUsc: destUsc,
	}
}
