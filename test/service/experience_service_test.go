package service

import (
	"fmt"
	"testing"

	"github.com/faridlan/firestore-go/model/web"
	experienceservice "github.com/faridlan/firestore-go/service/experience_service"
	"github.com/stretchr/testify/assert"
)

var experienceService experienceservice.ExperienceService

func TestExperienceFind(t *testing.T) {

	experienceResponse, err := experienceService.Find(ctx)
	assert.Nil(t, err)

	for _, experience := range experienceResponse {
		fmt.Println(experience)
	}

}

func TestExperienceFindId(t *testing.T) {

	experienceResponse, err := experienceService.FindId(ctx, "DR3RfWw4O1026rmKa3FD")
	assert.Nil(t, err)

	fmt.Println(experienceResponse)

}

func TestExperienceSave(t *testing.T) {

	experience := &web.ExperienceCreate{
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

	experienceResponse, err := experienceService.Save(ctx, experience)
	assert.Nil(t, err)

	fmt.Println(experienceResponse)

}

func TestExperienceDelete(t *testing.T) {

	experienceResponse, err := experienceService.FindId(ctx, "ArmcT7ZLCpzsDbjL3C96")
	assert.Nil(t, err)

	err = experienceService.Delete(ctx, experienceResponse.ID)
	assert.Nil(t, err)

	fmt.Println(experienceResponse)

}

func TestExperienceUpdate(t *testing.T) {

	experienceFind, err := experienceService.FindId(ctx, "LH6ONd0MDYUi2tFMHS4w")
	assert.Nil(t, err)

	experience := &web.ExperienceCreate{
		ID:          experienceFind.ID,
		CompanyName: "company_test_local_updated",
		Address:     "address_test_local_updated",
		Title:       "title_test_local_updated",
		EntryYear:   "2022",
		OutYear:     "present",
		JobDesk: []string{
			"job_desk_1_test_local_updated",
			"job_desk_2_test_local_updated",
		},
	}

	experienceResponse, err := experienceService.Update(ctx, experience)
	assert.Nil(t, err)

	fmt.Println(experienceResponse)

}
