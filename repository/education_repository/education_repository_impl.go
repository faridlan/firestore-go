package educationrepository

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/faridlan/firestore-go/model/domain"
	"google.golang.org/api/iterator"
)

type EducationRepositoryImpl struct {
}

func NewEducationRepository() EducationRepository {
	return &EducationRepositoryImpl{}
}

func (repository *EducationRepositoryImpl) Save(ctx context.Context, client *firestore.Client, education *domain.Education) (*domain.Education, error) {

	docRef, _, err := client.Collection("educations").Add(ctx, education)
	if err != nil {
		return nil, errors.New("failed to create new doc : " + err.Error())
	}

	education.ID = docRef.ID

	return education, err

}

func (repository *EducationRepositoryImpl) Update(ctx context.Context, client *firestore.Client, education *domain.Education) (*domain.Education, error) {

	docRef := client.Collection("educations").Doc(education.ID)

	updates := []firestore.Update{
		{
			Path:  "edu_name",
			Value: education.EduName,
		},
		{
			Path:  "address",
			Value: education.Address,
		},
		{
			Path:  "entry_year",
			Value: education.EntryYear,
		},
		{
			Path:  "out_year",
			Value: education.OutYear,
		},
		{
			Path:  "title",
			Value: education.Title,
		},
		{
			Path:  "achiev",
			Value: education.Achiev,
		},
	}

	_, err := docRef.Update(ctx, updates)
	if err != nil {
		return nil, errors.New("failed to update doc : " + err.Error())
	}

	education.ID = docRef.ID

	return education, nil
}

func (repository *EducationRepositoryImpl) Find(ctx context.Context, client *firestore.Client) ([]*domain.Education, error) {

	iter := client.Collection("educations").Documents(ctx)

	defer iter.Stop()

	var educations []*domain.Education

	for {
		docSnapshot, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, errors.New("failed to iterated : " + err.Error())
		}

		var education *domain.Education
		err = docSnapshot.DataTo(&education)
		if err != nil {
			return nil, errors.New("failed to decode doc : " + err.Error())
		}

		education.ID = docSnapshot.Ref.ID

		educations = append(educations, education)
	}

	return educations, nil

}

func (repository *EducationRepositoryImpl) FindId(ctx context.Context, client *firestore.Client, educationId string) (*domain.Education, error) {

	docRef := client.Collection("educations").Doc(educationId)

	docSnapshot, err := docRef.Get(ctx)
	if err != nil {
		return nil, errors.New("education not found")
	}

	var education *domain.Education
	docSnapshot.DataTo(&education)

	education.ID = docSnapshot.Ref.ID

	return education, nil

}

func (repository *EducationRepositoryImpl) Delete(ctx context.Context, client *firestore.Client, educationId string) error {

	docRef := client.Collection("educations").Doc(educationId)

	_, err := docRef.Delete(ctx)
	if err != nil {
		return errors.New("failed to delete doc : " + err.Error())
	}

	return nil
}
