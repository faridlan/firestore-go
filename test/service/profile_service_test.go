package service

import (
	"fmt"
	"testing"

	"github.com/faridlan/firestore-go/model/web"
	profileservice "github.com/faridlan/firestore-go/service/profile_service"
	"github.com/stretchr/testify/assert"
)

var profileService profileservice.ProfileService

func TestProfileFind(t *testing.T) {

	profileResponse, err := profileService.Find(ctx, "HxHIO0QGmdx0ACj7hYh8")
	assert.Nil(t, err)

	fmt.Println(profileResponse)

}

func TestProfileSave(t *testing.T) {

	profile := &web.Profile{
		Name:        "profile_service_test",
		Description: "description_service_test",
		Email:       "email_service@test.com",
		MediaSocial: web.MediaSocial{
			LinkedIn:  "linkedin.com/test",
			Instagram: "instagram.com/test",
			GitHub:    "github.com/test",
		},
		About: "about_service_test",
	}

	profileResponse, err := profileService.Save(ctx, profile)
	assert.Nil(t, err)

	fmt.Println(profileResponse)

}

func TestProfileDelete(t *testing.T) {

	profileResponse, err := profileService.Find(ctx, "QB0F8hLvpig8HBoo5K9l")
	assert.Nil(t, err)

	err = profileService.Delete(ctx, profileResponse.ID)
	assert.Nil(t, err)

}

func TestProfileUpdate(t *testing.T) {

	profileFind, err := profileService.Find(ctx, "gGoqJhnhqbzvXeXXKaoH")
	assert.Nil(t, err)

	profile := &web.Profile{
		ID:          profileFind.ID,
		Name:        "profile_service_test_updated",
		Description: "description_service_test_updated",
		Email:       "email_service@test.com",
		MediaSocial: web.MediaSocial{
			LinkedIn:  "linkedin.com/test_updated",
			Instagram: "instagram.com/test_updated",
			GitHub:    "github.com/test",
		},
		About: "about_service_test_updated",
	}

	profileResponse, err := profileService.Update(ctx, profile)
	assert.Nil(t, err)

	fmt.Println(profileResponse)

}
