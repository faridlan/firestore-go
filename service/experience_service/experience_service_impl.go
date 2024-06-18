package experienceservice

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/faridlan/firestore-go/exception"
	"github.com/faridlan/firestore-go/helper"
	"github.com/faridlan/firestore-go/model/domain"
	"github.com/faridlan/firestore-go/model/web"
	experiencerepository "github.com/faridlan/firestore-go/repository/experience_repository"
	"github.com/go-playground/validator/v10"
)

type ExperienceServiceImpl struct {
	ExperienceRepo experiencerepository.ExperienceRepository
	Client         *firestore.Client
	Validate       *validator.Validate
}

func NewExperienceService(exxperienceRepo experiencerepository.ExperienceRepository, client *firestore.Client, validate *validator.Validate) ExperienceService {

	return &ExperienceServiceImpl{
		ExperienceRepo: exxperienceRepo,
		Client:         client,
		Validate:       validate,
	}

}

func (service *ExperienceServiceImpl) Save(ctx context.Context, request *web.ExperienceCreate) (*web.ExperienceResponse, error) {

	err := service.Validate.Struct(request)
	errstring := helper.TranslateError(err, service.Validate)
	if err != nil {
		return nil, exception.NewBadRequestError(errstring)
	}

	experience := &domain.Experience{
		CompanyName: request.CompanyName,
		Address:     request.Address,
		Title:       request.Title,
		EntryYear:   request.EntryYear,
		OutYear:     request.OutYear,
		JobDesk:     request.JobDesk,
	}

	experienceResponse, err := service.ExperienceRepo.Save(ctx, service.Client, experience)
	if err != nil {
		return nil, err
	}

	return helper.ToExperienceResponse(experienceResponse), nil

}

func (service *ExperienceServiceImpl) Update(ctx context.Context, request *web.ExperienceCreate) (*web.ExperienceResponse, error) {

	err := service.Validate.Struct(request)
	errstring := helper.TranslateError(err, service.Validate)
	if err != nil {
		return nil, exception.NewBadRequestError(errstring)
	}

	experience, err := service.ExperienceRepo.FindId(ctx, service.Client, request.ID)
	if err != nil {
		return nil, err
	}

	experience.CompanyName = request.CompanyName
	experience.Address = request.Address
	experience.Title = request.Title
	experience.EntryYear = request.EntryYear
	experience.OutYear = request.OutYear
	experience.JobDesk = request.JobDesk

	experienceResponse, err := service.ExperienceRepo.Update(ctx, service.Client, experience)
	if err != nil {
		return nil, err
	}

	return helper.ToExperienceResponse(experienceResponse), nil

}

func (service *ExperienceServiceImpl) FindId(ctx context.Context, experienceId string) (*web.ExperienceResponse, error) {

	experienceResponse, err := service.ExperienceRepo.FindId(ctx, service.Client, experienceId)
	if err != nil {
		return nil, &exception.NotFound{
			Message: err.Error(),
		}
	}

	return helper.ToExperienceResponse(experienceResponse), nil

}

func (service *ExperienceServiceImpl) Find(ctx context.Context) ([]*web.ExperienceResponse, error) {

	experienceResponses, err := service.ExperienceRepo.Find(ctx, service.Client)
	if err != nil {
		return nil, nil
	}

	return helper.ToExperienceResponses(experienceResponses), nil

}

func (service *ExperienceServiceImpl) Delete(ctx context.Context, experienceId string) error {

	experience, err := service.ExperienceRepo.FindId(ctx, service.Client, experienceId)
	if err != nil {
		return &exception.NotFound{
			Message: err.Error(),
		}
	}

	err = service.ExperienceRepo.Delete(ctx, service.Client, experience.ID)
	if err != nil {
		return err
	}

	return nil

}
