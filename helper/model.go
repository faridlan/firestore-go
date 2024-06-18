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

func ToExperienceResponse(experience *domain.Experience) *web.ExperienceResponse {

	return &web.ExperienceResponse{
		ID:          experience.ID,
		CompanyName: experience.CompanyName,
		Address:     experience.Address,
		Title:       experience.Title,
		EntryYear:   experience.EntryYear,
		OutYear:     experience.OutYear,
		JobDesk:     experience.JobDesk,
	}

}

func ToExperienceResponses(experiences []*domain.Experience) []*web.ExperienceResponse {

	experienceResponses := []*web.ExperienceResponse{}
	for _, experience := range experiences {
		experienceResponses = append(experienceResponses, ToExperienceResponse(experience))
	}

	return experienceResponses

}
