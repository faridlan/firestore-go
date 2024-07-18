package controller

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/faridlan/firestore-go/config"
	educationcontroller "github.com/faridlan/firestore-go/controller/education_controller"
	experiencecontroller "github.com/faridlan/firestore-go/controller/experience_controller"
	profilecontroller "github.com/faridlan/firestore-go/controller/profile_controller"
	skillcontroller "github.com/faridlan/firestore-go/controller/skill_controller"
	educationrepository "github.com/faridlan/firestore-go/repository/education_repository"
	experiencerepository "github.com/faridlan/firestore-go/repository/experience_repository"
	profilerepository "github.com/faridlan/firestore-go/repository/profile_repository"
	skillrepository "github.com/faridlan/firestore-go/repository/skill_repository"
	educationservice "github.com/faridlan/firestore-go/service/education_service"
	experienceservice "github.com/faridlan/firestore-go/service/experience_service"
	profileservice "github.com/faridlan/firestore-go/service/profile_service"
	skillservice "github.com/faridlan/firestore-go/service/skill_service"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

var (
	skillController skillcontroller.SkillController
	client          *firestore.Client
	ctx             = context.Background()
)

func init() {
	if err := loadEnv(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	var err error
	client, err = config.NewDatabase()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	validate := validator.New()

	skillRepo := skillrepository.NewSkillRepository()
	skillService := skillservice.NewSkillService(skillRepo, client, validate)
	skillController = skillcontroller.NewSkillController(skillService)

	profileRepo := profilerepository.NewProfileRepository()
	profileService := profileservice.NewProfileService(profileRepo, client, validate)
	profileController = profilecontroller.NewProfileController(profileService)

	experienceRepo := experiencerepository.NewExperienceRepository()
	experienceService := experienceservice.NewExperienceService(experienceRepo, client, validate)
	experienceController = experiencecontroller.NewExperienceController(experienceService)

	educationRepo := educationrepository.NewEducationRepository()
	educationService := educationservice.NewEducationService(educationRepo, client, validate)
	educationController = educationcontroller.NewEducationController(educationService)

}

func loadEnv() error {
	// Specify the relative path to the .env file from the current directory
	err := godotenv.Load("../../.env")
	if err != nil {
		return err
	}
	return nil
}
