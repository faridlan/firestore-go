package repo

import (
	"fmt"
	"testing"

	"github.com/faridlan/firestore-go/model/domain"
	experiencerepository "github.com/faridlan/firestore-go/repository/experience_repository"
	"github.com/stretchr/testify/assert"
)

var experienceRepo experiencerepository.ExperienceRepository

func TestFindAllExperience(t *testing.T) {

	experienceResponse, err := experienceRepo.Find(ctx, client)
	assert.Nil(t, err)

	for _, experience := range experienceResponse {
		fmt.Println(experience)
	}

}

func TestFindByIdExperience(t *testing.T) {

	experienceResponse, err := experienceRepo.FindId(ctx, client, "Lu7SOjHFZkChYQxoUgVR")
	assert.Nil(t, err)

	fmt.Println(experienceResponse)

}

func TestSaveExperience(t *testing.T) {

	experience := &domain.Experience{
		CompanyName: "company_test_local",
		Address:     "address_test_local",
		Title:       "title_test_local",
		EntryYear:   "2022",
		OutYear:     "present",
		JobDesk: []string{
			"job_desk_1_test_local",
			"job_desk_2_test_local",
		},
	}

	experienceResponse, err := experienceRepo.Save(ctx, client, experience)
	assert.Nil(t, err)

	fmt.Println(experienceResponse)

}

func TestDeleteExperience(t *testing.T) {

	experience, err := experienceRepo.FindId(ctx, client, "ECrn6Et0IRoO3nF9iqif")
	assert.Nil(t, err)

	err = experienceRepo.Delete(ctx, client, experience.ID)
	assert.Nil(t, err)

}

func TestUpdateExperience(t *testing.T) {

	experience, err := experienceRepo.FindId(ctx, client, "YAt01EbwgYDFWZkketGK")
	assert.Nil(t, err)

	experience.CompanyName = "company_test_local_updated"
	experience.Address = "address_test_local_updated"
	experience.Title = "title_test_local_updated"
	experience.EntryYear = "2022"
	experience.OutYear = "present"
	experience.JobDesk = []string{
		"job_desk_1_test_local_updated",
		"job_desk_2_test_local_updated",
	}

	experienceResponse, err := experienceRepo.Update(ctx, client, experience)
	assert.Nil(t, err)

	fmt.Println(experienceResponse)

}
