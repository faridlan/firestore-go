package repo

import (
	"fmt"
	"testing"

	"github.com/faridlan/firestore-go/model/domain"
	profilerepository "github.com/faridlan/firestore-go/repository/profile_repository"
	"github.com/stretchr/testify/assert"
)

var repoProfile profilerepository.ProfileRepository

func TestGetProfile(t *testing.T) {

	profileResponse, err := repoProfile.Find(ctx, client, "HxHIO0QGmdx0ACj7hYh8")
	assert.Nil(t, err)
	assert.NotNil(t, profileResponse, "Expected profileResponse to be non-nil")
	t.Logf("Profile response: %+v", profileResponse)

}

func TestCreateProfile(t *testing.T) {

	profile := &domain.Profile{
		Name:        "name_test_local",
		Description: "description_test_local",
		Email:       "test_local@mail.com",
		MediaSocial: domain.MediaSocial{
			LinkedIn:  "linkedin/user_test",
			Instagram: "instagram/user_test",
			GitHub:    "github.com/user_test",
		},
		About: "about_test_local",
	}

	profileResponse, err := repoProfile.Save(ctx, client, profile)
	assert.Nil(t, err)

	fmt.Println(profileResponse)

}

func TestDeleteProfile(t *testing.T) {

	profileResponse, err := repoProfile.Find(ctx, client, "vSQSGeJNFdH1C7GBwu59")
	assert.Nil(t, err)

	err = repoProfile.Delete(ctx, client, profileResponse.ID)
	assert.Nil(t, err)
}

func TestUpdateProfile(t *testing.T) {
	profileResponse, err := repoProfile.Find(ctx, client, "WDOFjpiqPm4CP6FlYMjQ")
	assert.Nil(t, err)

	profileResponse.Name = "user_update"
	profileResponse.Email = "email_update@mail.com"
	profileResponse.Description = "description test"
	profileResponse.MediaSocial.GitHub = "github.com/update"
	profileResponse.MediaSocial.Instagram = "instagram.com/update"
	profileResponse.MediaSocial.LinkedIn = "linkedin.com/update"
	profileResponse.About = "about update"

	respone, err := repoProfile.Update(ctx, client, profileResponse)
	assert.Nil(t, err)

	fmt.Println(respone)
}
