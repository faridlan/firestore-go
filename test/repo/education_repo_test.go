package repo

import (
	"fmt"
	"testing"

	"github.com/faridlan/firestore-go/model/domain"
	educationrepository "github.com/faridlan/firestore-go/repository/education_repository"
	"github.com/stretchr/testify/assert"
)

var educationRepo educationrepository.EducationRepository

func TestEducationFind(t *testing.T) {

	educationResponse, err := educationRepo.Find(ctx, client)
	assert.Nil(t, err)

	for _, education := range educationResponse {
		fmt.Println(education)
	}

}

func TestEducationFindId(t *testing.T) {

	educationResponse, err := educationRepo.FindId(ctx, client, "Kp0LEFQc3DiMjp7stqgk")
	assert.Nil(t, err)

	fmt.Println(educationResponse)

}

func TestEducationSave(t *testing.T) {

	education := &domain.Education{
		EduName:   "edu_test_local",
		Address:   "address_test_local",
		EntryYear: "2018",
		OutYear:   "2022",
		Title:     "bachelot_test_local",
		Achiev: []string{
			"create project local test",
			"create project local test 2",
		},
	}

	educationResponse, err := educationRepo.Save(ctx, client, education)
	assert.Nil(t, err)

	fmt.Println(educationResponse)

}

func TestEducationDelete(t *testing.T) {

	educationResponse, err := educationRepo.FindId(ctx, client, "6IYagzhtvkkz2SdlNuR4")
	assert.Nil(t, err)

	err = educationRepo.Delete(ctx, client, educationResponse.ID)
	assert.Nil(t, err)

}

func TestEducationUpdate(t *testing.T) {

	education, err := educationRepo.FindId(ctx, client, "o90UsxkgiGLbXftrzjlr")
	assert.Nil(t, err)

	education.EduName = "edu_test_local_updated"
	education.Address = "address_test_local_updated"
	education.EntryYear = "2018"
	education.OutYear = "2022"
	education.Title = "bachelot_test_local_updated"
	education.Achiev = []string{
		"create project local test_updated",
		"create project local test 2_updated",
	}

	educationResponse, err := educationRepo.Update(ctx, client, education)
	assert.Nil(t, err)

	fmt.Println(educationResponse)

}
