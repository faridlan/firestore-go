package profileservice

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/faridlan/firestore-go/model/web"
	profilerepository "github.com/faridlan/firestore-go/repository/profile_repository"
	"github.com/go-playground/validator/v10"
)

type ProfileServiceimpl struct {
	ProfileRepo profilerepository.ProfileRepository
	Client      *firestore.Client
	Validate    *validator.Validate
}

func NewProfileController(profileRepo profilerepository.ProfileRepository, client *firestore.Client, validate *validator.Validate) ProfileService {
	return &ProfileServiceimpl{
		ProfileRepo: profileRepo,
		Client:      client,
		Validate:    validate,
	}
}

func (service *ProfileServiceimpl) Save(ctx context.Context, request *web.Profile) (*web.Profile, error) {
	panic("not implemented") // TODO: Implement
}

func (service *ProfileServiceimpl) find(ctx context.Context, profileId string) (*web.Profile, error) {
	panic("not implemented") // TODO: Implement
}

func (service *ProfileServiceimpl) Update(ctx context.Context, profileId string) (*web.Profile, error) {
	panic("not implemented") // TODO: Implement
}

func (service *ProfileServiceimpl) Delete(ctx context.Context, profileId string) error {
	panic("not implemented") // TODO: Implement
}
