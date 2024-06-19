package experiencerepository

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/faridlan/firestore-go/model/domain"
	"google.golang.org/api/iterator"
)

type ExperienceRepositoryImpl struct {
}

func NewExperienceRepository() ExperienceRepository {
	return &ExperienceRepositoryImpl{}
}

func (repository *ExperienceRepositoryImpl) Save(ctx context.Context, client *firestore.Client, experience *domain.Experience) (*domain.Experience, error) {

	docRef, _, err := client.Collection("experiences").Add(ctx, experience)
	if err != nil {
		return nil, errors.New("failed to add new document : " + err.Error())
	}

	experience.ID = docRef.ID

	return experience, nil

}

func (repository *ExperienceRepositoryImpl) Update(ctx context.Context, client *firestore.Client, experience *domain.Experience) (*domain.Experience, error) {

	docRef := client.Collection("experiences").Doc(experience.ID)

	updates := []firestore.Update{
		{
			Path:  "company_name",
			Value: experience.CompanyName,
		},
		{
			Path:  "address",
			Value: experience.Address,
		},
		{
			Path:  "title",
			Value: experience.Title,
		},
		{
			Path:  "entry_year",
			Value: experience.EntryYear,
		},
		{
			Path:  "out_year",
			Value: experience.OutYear,
		},
		{
			Path:  "job_desk",
			Value: experience.JobDesk,
		},
	}

	_, err := docRef.Update(ctx, updates)
	if err != nil {
		return nil, errors.New("failed to update document : " + err.Error())
	}

	experience.ID = docRef.ID

	return experience, nil

}

func (repository *ExperienceRepositoryImpl) Find(ctx context.Context, client *firestore.Client) ([]*domain.Experience, error) {

	iter := client.Collection("experiences").Documents(ctx)

	defer iter.Stop()

	clients := []*domain.Experience{}

	for {
		docSnapshot, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, errors.New("failed to iterate through documents : " + err.Error())
		}

		client := domain.Experience{}
		err = docSnapshot.DataTo(&client)
		if err != nil {
			return nil, errors.New("failed to decode documents : " + err.Error())
		}

		clients = append(clients, &client)
	}

	return clients, nil

}

func (repository *ExperienceRepositoryImpl) FindId(ctx context.Context, client *firestore.Client, experienceId string) (*domain.Experience, error) {

	docRef := client.Collection("experiences").Doc(experienceId)

	docSnapshot, err := docRef.Get(ctx)
	if err != nil {
		return nil, errors.New("experience not found")
	}

	var experience domain.Experience
	err = docSnapshot.DataTo(&experience)
	if err != nil {
		return nil, errors.New("failed to encode documents : " + err.Error())
	}

	experience.ID = docSnapshot.Ref.ID

	return &experience, nil

}

func (repository *ExperienceRepositoryImpl) Delete(ctx context.Context, client *firestore.Client, experienceId string) error {

	docRef := client.Collection("experiences").Doc(experienceId)

	_, err := docRef.Delete(ctx)
	if err != nil {
		return errors.New("failed to delete experience : " + err.Error())
	}

	return nil

}
