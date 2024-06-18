package experiencerepository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/faridlan/firestore-go/model/domain"
)

type ExperienceRepository interface {
	Save(ctx context.Context, client *firestore.Client, experience *domain.Experience) (*domain.Experience, error)
	Update(ctx context.Context, client *firestore.Client, experience *domain.Experience) (*domain.Experience, error)
	Find(ctx context.Context, client *firestore.Client) ([]*domain.Experience, error)
	FindId(ctx context.Context, client *firestore.Client, experienceId string) (*domain.Experience, error)
	Delete(ctx context.Context, client *firestore.Client, experienceId string) error
}
