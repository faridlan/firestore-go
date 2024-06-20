package educationservice

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/faridlan/firestore-go/exception"
	"github.com/faridlan/firestore-go/helper"
	"github.com/faridlan/firestore-go/model/domain"
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

	err := service.Validate.Struct(request)
	errString := helper.TranslateError(err, service.Validate)
	if err != nil {
		return nil, exception.NewBadRequestError(errString)
	}

	education := &domain.Education{
		EduName:   request.EduName,
		Address:   request.Address,
		EntryYear: request.EntryYear,
		OutYear:   request.OutYear,
		Title:     request.Title,
		Achiev:    request.Achiev,
	}

	educationResponse, err := service.EducationRepo.Save(ctx, service.Client, education)
	if err != nil {
		return nil, err
	}

	return helper.ToEducationResponse(educationResponse), nil

}

func (service *EducationServiceImpl) Update(ctx context.Context, request *web.EducationCreate) (*web.EducationResponse, error) {

	err := service.Validate.Struct(request)
	errString := helper.TranslateError(err, service.Validate)
	if err != nil {
		return nil, exception.NewBadRequestError(errString)
	}

	education, err := service.EducationRepo.FindId(ctx, service.Client, request.ID)
	if err != nil {
		return nil, &exception.NotFound{
			Message: err.Error(),
		}
	}

	education.EduName = request.EduName
	education.Address = request.Address
	education.EntryYear = request.EntryYear
	education.OutYear = request.OutYear
	education.Title = request.Title
	education.Achiev = request.Achiev

	educationResponse, err := service.EducationRepo.Update(ctx, service.Client, education)
	if err != nil {
		return nil, err
	}

	return helper.ToEducationResponse(educationResponse), nil

}

func (service *EducationServiceImpl) Find(ctx context.Context) ([]*web.EducationResponse, error) {

	educationResponses, err := service.EducationRepo.Find(ctx, service.Client)
	if err != nil {
		return nil, err
	}

	return helper.ToEducationResponses(educationResponses), nil

}

func (service *EducationServiceImpl) FindId(ctx context.Context, educationId string) (*web.EducationResponse, error) {

	educationResponse, err := service.EducationRepo.FindId(ctx, service.Client, educationId)
	if err != nil {
		return nil, &exception.NotFound{
			Message: err.Error(),
		}
	}

	return helper.ToEducationResponse(educationResponse), nil

}

func (service *EducationServiceImpl) Delete(ctx context.Context, educationId string) error {

	education, err := service.EducationRepo.FindId(ctx, service.Client, educationId)
	if err != nil {
		return &exception.NotFound{
			Message: err.Error(),
		}
	}

	service.EducationRepo.Delete(ctx, service.Client, education.ID)

	return nil

}
