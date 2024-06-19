package educationrepository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/faridlan/firestore-go/model/domain"
)

type EducationRepository interface {
	Save(ctx context.Context, client *firestore.Client, education *domain.Education) (*domain.Education, error)
	Update(ctx context.Context, client *firestore.Client, education *domain.Education) (*domain.Education, error)
	Find(ctx context.Context, client *firestore.Client) ([]*domain.Education, error)
	FindId(ctx context.Context, client *firestore.Client, educationId string) (*domain.Education, error)
	Delete(ctx context.Context, client *firestore.Client, educationId string) error
}
