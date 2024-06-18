package main

import (
	"github.com/faridlan/firestore-go/config"
	experiencecontroller "github.com/faridlan/firestore-go/controller/experience_controller"
	profilecontroller "github.com/faridlan/firestore-go/controller/profile_controller"
	experiencerepository "github.com/faridlan/firestore-go/repository/experience_repository"
	profilerepository "github.com/faridlan/firestore-go/repository/profile_repository"
	experienceservice "github.com/faridlan/firestore-go/service/experience_service"
	profileservice "github.com/faridlan/firestore-go/service/profile_service"
	"github.com/go-playground/validator/v10"
)

func main() {

	db, _ := config.NewDatabase()
	validate := validator.New()

	profileRepo := profilerepository.NewProfileRepository()
	profileService := profileservice.NewProfileService(profileRepo, db, validate)
	profileController := profilecontroller.NewProfileController(profileService)

	experienceRepo := experiencerepository.NewExperienceRepository()
	experienceService := experienceservice.NewExperienceService(experienceRepo, db, validate)
	experienceController := experiencecontroller.NewExperienceController(experienceService)

	routers := config.Router{
		ProfileContrroler:    profileController,
		ExperienceContrroler: experienceController,
	}

	app := config.NewRouter(routers)

	app.Listen("localhost:9090")

}
