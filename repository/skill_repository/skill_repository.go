package skillrepository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/faridlan/firestore-go/model/domain"
)

type SkillRepository interface {
	Save(ctx context.Context, client *firestore.Client, skill *domain.Skill) (*domain.Skill, error)
	Update(ctx context.Context, client *firestore.Client, skill *domain.Skill) (*domain.Skill, error)
	Find(ctx context.Context, client *firestore.Client) ([]*domain.Skill, error)
	FindId(ctx context.Context, client *firestore.Client) (*domain.Skill, error)
	Delete(ctx context.Context, client *firestore.Client, skillId string) error
}
