package skillrepository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/faridlan/firestore-go/model/domain"
)

type SkillRepositoryImpl struct {
}

func NewSkillRepository() SkillRepository {
	return &SkillRepositoryImpl{}
}

func (repository *SkillRepositoryImpl) Save(ctx context.Context, client *firestore.Client, skill *domain.Skill) (*domain.Skill, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *SkillRepositoryImpl) Update(ctx context.Context, client *firestore.Client, skill *domain.Skill) (*domain.Skill, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *SkillRepositoryImpl) Find(ctx context.Context, client *firestore.Client) ([]*domain.Skill, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *SkillRepositoryImpl) FindId(ctx context.Context, client *firestore.Client) (*domain.Skill, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *SkillRepositoryImpl) Delete(ctx context.Context, client *firestore.Client, skillId string) error {
	panic("not implemented") // TODO: Implement
}
