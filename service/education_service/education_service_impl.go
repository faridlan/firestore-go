package educationservice

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/faridlan/firestore-go/model/web"
	educationrepository "github.com/faridlan/firestore-go/repository/education_repository"
	"github.com/go-playground/validator/v10"
)

type EducationServiceImpl struct {
	EducationRepo educationrepository.EducationRepository
	Client        *firestore.Client
	Validate      *validator.Validate
}

func NewEducationService(educationRepo educationrepository.EducationRepository, client *firestore.Client, validate *validator.Validate) EducationService {
	return &EducationServiceImpl{
		EducationRepo: educationRepo,
		Client:        client,
		Validate:      validate,
	}
}

func (service *EducationServiceImpl) Save(ctx context.Context, request *web.EducationCreate) (*web.EducationResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (service *EducationServiceImpl) Update(ctx context.Context, request *web.EducationCreate) (*web.EducationResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (service *EducationServiceImpl) Find(ctx context.Context) ([]*web.EducationResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (service *EducationServiceImpl) FindId(ctx context.Context, educationId string) (*web.EducationResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (service *EducationServiceImpl) Delete(ctx context.Context, educationId string) error {
	panic("not implemented") // TODO: Implement
}
