package repo

import (
	"context"
	"fmt"
	"log"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/faridlan/firestore-go/config"
	"github.com/faridlan/firestore-go/model/domain"
	profilerepository "github.com/faridlan/firestore-go/repository/profile_repository"
	skillrepository "github.com/faridlan/firestore-go/repository/skill_repository"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var (
	client    *firestore.Client
	skillRepo skillrepository.SkillRepository
	ctx       = context.Background()
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

	skillRepo = skillrepository.NewSkillRepository()
	repoProfile = profilerepository.NewProfileRepository()
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
	skillResponse, err := skillRepo.Find(ctx, client)
	assert.Nil(t, err)

	for _, skill := range skillResponse {
		fmt.Println(skill.ID)
		fmt.Println(skill.Name)
	}
}

func TestSkillFindId(t *testing.T) {

	skillResponse, err := skillRepo.FindId(ctx, client, "3nBpOEmlcIJVFEbQNS00")
	assert.Nil(t, err)

	fmt.Println(skillResponse)
}

func TestSkillCreate(t *testing.T) {

	skill := &domain.Skill{
		Name: "skill-unit-test-again",
	}

	skillResponse, err := skillRepo.Save(ctx, client, skill)
	assert.Nil(t, err)

	fmt.Println(skillResponse)
}

func TestSkillDelete(t *testing.T) {

	err := skillRepo.Delete(ctx, client, "3nBpOEmlcIJVFEbQNS00")
	assert.Nil(t, err)

}

func TestSkillUpdate(t *testing.T) {

	skill, err := skillRepo.FindId(ctx, client, "1AHCvPivE1VT2xikG4NP")
	assert.Nil(t, err)

	skill.Name = "skill-test-update"

	skillResponse, err := skillRepo.Update(ctx, client, skill)
	assert.Nil(t, err)

	fmt.Println(skillResponse)

	assert.Equal(t, skill.Name, skillResponse.Name)
}
