package main

import (
	"log"

	"github.com/faridlan/firestore-go/config"
	educationcontroller "github.com/faridlan/firestore-go/controller/education_controller"
	experiencecontroller "github.com/faridlan/firestore-go/controller/experience_controller"
	profilecontroller "github.com/faridlan/firestore-go/controller/profile_controller"
	skillcontroller "github.com/faridlan/firestore-go/controller/skill_controller"
	"github.com/faridlan/firestore-go/helper"
	educationrepository "github.com/faridlan/firestore-go/repository/education_repository"
	experiencerepository "github.com/faridlan/firestore-go/repository/experience_repository"
	profilerepository "github.com/faridlan/firestore-go/repository/profile_repository"
	skillrepository "github.com/faridlan/firestore-go/repository/skill_repository"
	educationservice "github.com/faridlan/firestore-go/service/education_service"
	experienceservice "github.com/faridlan/firestore-go/service/experience_service"
	profileservice "github.com/faridlan/firestore-go/service/profile_service"
	skillservice "github.com/faridlan/firestore-go/service/skill_service"
	"github.com/go-playground/validator/v10"
)

func main() {

	err := helper.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	db, _ := config.NewDatabase()
	validate := validator.New()

	profileRepo := profilerepository.NewProfileRepository()
	profileService := profileservice.NewProfileService(profileRepo, db, validate)
	profileController := profilecontroller.NewProfileController(profileService)

	experienceRepo := experiencerepository.NewExperienceRepository()
	experienceService := experienceservice.NewExperienceService(experienceRepo, db, validate)
	experienceController := experiencecontroller.NewExperienceController(experienceService)

	educationRepo := educationrepository.NewEducationRepository()
	educationService := educationservice.NewEducationService(educationRepo, db, validate)
	educationController := educationcontroller.NewEducationController(educationService)

	skillrepo := skillrepository.NewSkillRepository()
	skillService := skillservice.NewSkillService(skillrepo, db, validate)
	skillController := skillcontroller.NewSkillController(skillService)

	routers := config.Router{
		ProfileContrroler:    profileController,
		ExperienceContrroler: experienceController,
		EducationController:  educationController,
		SkillController:      skillController,
	}

	app := config.NewRouter(routers)

	app.Listen("localhost:9090")

}
