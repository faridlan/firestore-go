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

func ToEducationResponse(education *domain.Education) *web.EducationResponse {

	return &web.EducationResponse{
		ID:        education.ID,
		EduName:   education.EduName,
		Address:   education.Address,
		EntryYear: education.EntryYear,
		OutYear:   education.OutYear,
		Title:     education.Title,
		Achiev:    education.Achiev,
	}

}

func ToEducationResponses(educations []*domain.Education) []*web.EducationResponse {

	educactionResponses := []*web.EducationResponse{}

	for _, education := range educations {
		educactionResponses = append(educactionResponses, ToEducationResponse(education))
	}

	return educactionResponses

}

func ToSkillResposne(skill *domain.Skill) *web.SkillResponse {

	return &web.SkillResponse{
		ID:   skill.ID,
		Name: skill.Name,
	}

}

func ToSkillResposnes(skills []*domain.Skill) []*web.SkillResponse {

	skillResponses := []*web.SkillResponse{}

	for _, skill := range skills {
		skillResponses = append(skillResponses, ToSkillResposne(skill))
	}

	return skillResponses

}
