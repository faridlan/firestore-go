package service

import (
	"context"
	"fmt"
	"log"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/faridlan/firestore-go/config"
	"github.com/faridlan/firestore-go/model/web"
	skillrepository "github.com/faridlan/firestore-go/repository/skill_repository"
	skillservice "github.com/faridlan/firestore-go/service/skill_service"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var (
	skillService skillservice.SkillService
	client       *firestore.Client
	ctx          = context.Background()
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
	skillService = skillservice.NewSkillService(skillRepo, client, validate)

}

func loadEnv() error {
	// Specify the relative path to the .env file from the current directory
	err := godotenv.Load("../../.env")
	if err != nil {
		return err
	}
	return nil
}

func TestSkillFind(t *testing.T) {

	skillResponse, err := skillService.Find(ctx)
	assert.Nil(t, err)

	for _, skill := range skillResponse {
		fmt.Println(skill)
	}

}

func TestSkillFindId(t *testing.T) {

	skillResponse, err := skillService.FindId(ctx, "bodX9IS7s4sN3E6Di3xG")
	assert.Nil(t, err)

	fmt.Println(skillResponse)

}

func TestSkillSave(t *testing.T) {

	skill := &web.SkillCreate{
		Name: "skill_service_test_local",
	}

	skillResponse, err := skillService.Save(ctx, skill)
	assert.Nil(t, err)

	fmt.Println(skillResponse)

}

func TestSkillDelete(t *testing.T) {

	skillResponse, err := skillService.FindId(ctx, "bodX9IS7s4sN3E6Di3xG")
	assert.Nil(t, err)

	err = skillService.Delete(ctx, skillResponse.ID)
	assert.Nil(t, err)

}

func TestSkillUpdate(t *testing.T) {

	_, err := skillService.FindId(ctx, "lgD3Bo8RjncOY31G8woj")
	assert.Nil(t, err)

	skill := &web.SkillCreate{
		ID:   "lgD3Bo8RjncOY31G8woj",
		Name: "skill-2-updated",
	}

	skillResponse, err := skillService.Update(ctx, skill)
	assert.Nil(t, err)

	fmt.Println(skillResponse)

}
