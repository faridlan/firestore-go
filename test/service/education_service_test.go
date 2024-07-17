package service

import (
	"fmt"
	"testing"

	"github.com/faridlan/firestore-go/model/web"
	educationservice "github.com/faridlan/firestore-go/service/education_service"
	"github.com/stretchr/testify/assert"
)

var educationService educationservice.EducationService

func TestEducationFind(t *testing.T) {

	educationResponse, err := educationService.Find(ctx)
	assert.Nil(t, err)

	for _, education := range educationResponse {
		fmt.Println(education)
	}

}

func TestEducationFindId(t *testing.T) {

	educationResponse, err := educationService.FindId(ctx, "Kp0LEFQc3DiMjp7stqgk")
	assert.Nil(t, err)

	fmt.Println(educationResponse)

}

func TestEducationSave(t *testing.T) {

	education := &web.EducationCreate{
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

	educationResponse, err := educationService.Save(ctx, education)
	assert.Nil(t, err)

	fmt.Println(educationResponse)

}

func TestEducationDelete(t *testing.T) {

	educationResponse, err := educationService.FindId(ctx, "J2jTs1iBJA2DIUR4AynQ")
	assert.Nil(t, err)

	err = educationService.Delete(ctx, educationResponse.ID)
	assert.Nil(t, err)

}

func TestEducationUpdate(t *testing.T) {

	educationFind, err := educationService.FindId(ctx, "IRD9GHXF8BF8NpE3ATPG")
	assert.Nil(t, err)

	education := &web.EducationCreate{
		ID:        educationFind.ID,
		EduName:   "edu_test_local_updated",
		Address:   "address_test_local_updated",
		EntryYear: "2018",
		OutYear:   "2022",
		Title:     "bachelot_test_local_updated",
		Achiev: []string{
			"create project local test_updated",
			"create project local test 2_updated",
		},
	}

	educationResponse, err := educationService.Update(ctx, education)
	assert.Nil(t, err)

	fmt.Println(educationResponse)
}
