package profilerepository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/faridlan/firestore-go/model/domain"
)

type ProfileRepositoryImpl struct {
}

func NewProfileRepository() ProfileRepository {
	return &ProfileRepositoryImpl{}
}

func (repository *ProfileRepositoryImpl) Save(ctx context.Context, client firestore.Client, profile *domain.Profile) (*domain.Profile, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *ProfileRepositoryImpl) Find(ctx context.Context, client firestore.Client, profileId string) (*domain.Profile, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *ProfileRepositoryImpl) Update(ctx context.Context, client firestore.Client, profileId string) (*domain.Profile, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *ProfileRepositoryImpl) Delete(ctx context.Context, client firestore.Client, profileId string) error {
	panic("not implemented") // TODO: Implement
}
