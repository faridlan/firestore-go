package main

import (
	"github.com/faridlan/firestore-go/config"
	profilecontroller "github.com/faridlan/firestore-go/controller/profile_controller"
	profilerepository "github.com/faridlan/firestore-go/repository/profile_repository"
	profileservice "github.com/faridlan/firestore-go/service/profile_service"
	"github.com/go-playground/validator/v10"
)

func main() {

	db, _ := config.NewDatabase()
	validate := validator.New()

	profileRepo := profilerepository.NewProfileRepository()
	profileService := profileservice.NewProfileService(profileRepo, db, validate)
	profileController := profilecontroller.NewProfileController(profileService)

	routers := config.Router{
		ProfileContrroler: profileController,
	}

	app := config.NewRouter(routers)

	app.Listen("localhost:9090")

}
