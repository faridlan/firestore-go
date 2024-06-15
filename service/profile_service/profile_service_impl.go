package profileservice

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/faridlan/firestore-go/exception"
	"github.com/faridlan/firestore-go/helper"
	"github.com/faridlan/firestore-go/model/domain"
	"github.com/faridlan/firestore-go/model/web"
	profilerepository "github.com/faridlan/firestore-go/repository/profile_repository"
	"github.com/go-playground/validator/v10"
)

type ProfileServiceimpl struct {
	ProfileRepo profilerepository.ProfileRepository
	Client      *firestore.Client
	Validate    *validator.Validate
}

func NewProfileService(profileRepo profilerepository.ProfileRepository, client *firestore.Client, validate *validator.Validate) ProfileService {
	return &ProfileServiceimpl{
		ProfileRepo: profileRepo,
		Client:      client,
		Validate:    validate,
	}
}

func (service *ProfileServiceimpl) Save(ctx context.Context, request *web.Profile) (*web.Profile, error) {

	err := service.Validate.Struct(request)
	errString := helper.TranslateError(err, service.Validate)
	if err != nil {
		return nil, exception.NewBadRequestError(errString)
	}

	profile := &domain.Profile{
		Name:        request.Name,
		Description: request.Description,
		Email:       request.Email,
		MediaSocial: domain.MediaSocial{
			LinkedIn:  request.MediaSocial.LinkedIn,
			Instagram: request.MediaSocial.Instagram,
			GitHub:    request.MediaSocial.GitHub,
		},
		About: request.About,
	}

	profileResponse, err := service.ProfileRepo.Save(ctx, service.Client, profile)
	if err != nil {
		return nil, err
	}

	return helper.ToProfileResponse(profileResponse), nil

}

func (service *ProfileServiceimpl) Find(ctx context.Context, profileId string) (*web.Profile, error) {

	profileResponse, err := service.ProfileRepo.Find(ctx, service.Client, profileId)
	if err != nil {
		return nil, &exception.NotFound{
			Message: err.Error(),
		}
	}

	return helper.ToProfileResponse(profileResponse), nil

}

func (service *ProfileServiceimpl) Update(ctx context.Context, profileId string) (*web.Profile, error) {
	panic("not implemented") // TODO: Implement
}

func (service *ProfileServiceimpl) Delete(ctx context.Context, profileId string) error {
	panic("not implemented") // TODO: Implement
}
