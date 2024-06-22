package skillservice

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/faridlan/firestore-go/exception"
	"github.com/faridlan/firestore-go/helper"
	"github.com/faridlan/firestore-go/model/domain"
	"github.com/faridlan/firestore-go/model/web"
	skillrepository "github.com/faridlan/firestore-go/repository/skill_repository"
	"github.com/go-playground/validator/v10"
)

type SkillServiceImpl struct {
	SkillRepo skillrepository.SkillRepository
	Client    *firestore.Client
	Validate  *validator.Validate
}

func NewSkillService(skillRepo skillrepository.SkillRepository, client *firestore.Client, validate *validator.Validate) SkillService {
	return &SkillServiceImpl{
		SkillRepo: skillRepo,
		Client:    client,
		Validate:  validate,
	}
}

func (service *SkillServiceImpl) Save(ctx context.Context, request *web.SkillCreate) (*web.SkillResponse, error) {

	err := service.Validate.Struct(request)
	errString := helper.TranslateError(err, service.Validate)
	if err != nil {
		return nil, exception.NewBadRequestError(errString)
	}

	skill := &domain.Skill{
		Name: request.Name,
	}

	skillResponse, err := service.SkillRepo.Save(ctx, service.Client, skill)
	if err != nil {
		return nil, err
	}

	return helper.ToSkillResposne(skillResponse), nil

}

func (service *SkillServiceImpl) Update(ctx context.Context, request *web.SkillCreate) (*web.SkillResponse, error) {

	err := service.Validate.Struct(request)
	errString := helper.TranslateError(err, service.Validate)
	if err != nil {
		return nil, exception.NewBadRequestError(errString)
	}

	skill, err := service.SkillRepo.FindId(ctx, service.Client, request.ID)
	if err != nil {
		return nil, &exception.NotFound{
			Message: err.Error(),
		}
	}

	skill.Name = request.Name

	skillResposne, err := service.SkillRepo.Update(ctx, service.Client, skill)
	if err != nil {
		return nil, err
	}

	return helper.ToSkillResposne(skillResposne), nil

}

func (service *SkillServiceImpl) Find(ctx context.Context) ([]*web.SkillResponse, error) {

	skillResponses, err := service.SkillRepo.Find(ctx, service.Client)
	if err != nil {
		return nil, err
	}

	return helper.ToSkillResposnes(skillResponses), nil

}

func (service *SkillServiceImpl) FindId(ctx context.Context, SkillId string) (*web.SkillResponse, error) {

	skillResposne, err := service.SkillRepo.FindId(ctx, service.Client, SkillId)
	if err != nil {
		return nil, &exception.NotFound{
			Message: err.Error(),
		}
	}

	return helper.ToSkillResposne(skillResposne), nil

}

func (service *SkillServiceImpl) Delete(ctx context.Context, SkillId string) error {

	skill, err := service.SkillRepo.FindId(ctx, service.Client, SkillId)
	if err != nil {
		return &exception.NotFound{
			Message: err.Error(),
		}
	}

	err = service.SkillRepo.Delete(ctx, service.Client, skill.ID)
	if err != nil {
		return err
	}

	return nil

}
