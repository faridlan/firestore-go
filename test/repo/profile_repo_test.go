package repo

import (
	"fmt"
	"testing"

	"github.com/faridlan/firestore-go/config"
	"github.com/faridlan/firestore-go/model/domain"
	"github.com/stretchr/testify/assert"
)

var repoProfile profilerepository.ProfileRepository


func TestGetProfile(t *testing.T) {

	clients, err := config.NewDatabase()
	assert.Nil(t, err)

	profileResponse, err := repoProfile.Find(ctx, clients, "HxHIO0QGmdx0ACj7hYh8")
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
