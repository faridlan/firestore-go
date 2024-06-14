package helper

import (
	"github.com/faridlan/firestore-go/model/domain"
	"github.com/faridlan/firestore-go/model/web"
)

func ToProfileResponse(profile *domain.Profile) *web.Profile {

	return &web.Profile{
		ID:          profile.ID,
		Name:        profile.Name,
		Description: profile.Description,
		Email:       profile.Email,
		MediaSocial: web.MediaSocial{
			LinkedIn:  profile.MediaSocial.LinkedIn,
			Instagram: profile.MediaSocial.Instagram,
			GitHub:    profile.MediaSocial.GitHub,
		},
		About: profile.About,
	}

}
