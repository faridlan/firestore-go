package skillservice

import (
	"context"

	"cloud.google.com/go/firestore"
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
	panic("not implemented") // TODO: Implement
}

func (service *SkillServiceImpl) Update(ctx context.Context, request *web.SkillCreate) (*web.SkillResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (service *SkillServiceImpl) Find(ctx context.Context) ([]*web.SkillResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (service *SkillServiceImpl) FindId(ctx context.Context, SkillId string) (*web.SkillResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (service *SkillServiceImpl) Delete(ctx context.Context, SkillId string) error {
	panic("not implemented") // TODO: Implement
}
