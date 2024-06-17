package profilerepository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/faridlan/firestore-go/model/domain"
)

type ProfileRepository interface {
	Save(ctx context.Context, client *firestore.Client, profile *domain.Profile) (*domain.Profile, error)
	Find(ctx context.Context, client *firestore.Client, profileId string) (*domain.Profile, error)
	Update(ctx context.Context, client *firestore.Client, profile *domain.Profile) (*domain.Profile, error)
	Delete(ctx context.Context, client *firestore.Client, profileId string) error
}
