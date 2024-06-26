package profilerepository

import (
	"context"
	"errors"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/faridlan/firestore-go/model/domain"
)

type ProfileRepositoryImpl struct {
}

func NewProfileRepository() ProfileRepository {
	return &ProfileRepositoryImpl{}
}

func (repository *ProfileRepositoryImpl) Save(ctx context.Context, client *firestore.Client, profile *domain.Profile) (*domain.Profile, error) {

	// Check if the client is nil
	if client == nil {
		return nil, fmt.Errorf("firestore client is nil")
	}

	// Check if the profile is nil
	if profile == nil {
		return nil, fmt.Errorf("profile is nil")
	}

	docRef, _, err := client.Collection("profiles").Add(ctx, profile)
	if err != nil {
		return nil, fmt.Errorf("failed to add a new document: %v", err)
	}

	profile.ID = docRef.ID

	return profile, nil

}

func (repository *ProfileRepositoryImpl) Find(ctx context.Context, client *firestore.Client, profileId string) (*domain.Profile, error) {

	doc := client.Collection("profiles").Doc(profileId)

	docSnapshot, err := doc.Get(ctx)
	if err != nil {
		return nil, errors.New("profile not found")
	}

	var profile domain.Profile

	err = docSnapshot.DataTo(&profile)
	if err != nil {
		return nil, errors.New("failed to decode collection :" + err.Error())
	}

	profile.ID = docSnapshot.Ref.ID

	return &profile, nil

}

func (repository *ProfileRepositoryImpl) Update(ctx context.Context, client *firestore.Client, profile *domain.Profile) (*domain.Profile, error) {

	docRef := client.Collection("profiles").Doc(profile.ID)

	updates := []firestore.Update{
		{
			Path:  "name",
			Value: profile.Name,
		},
		{
			Path:  "description",
			Value: profile.Description,
		},
		{
			Path:  "email",
			Value: profile.Email,
		},
		{
			Path:  "media_social",
			Value: profile.MediaSocial,
		},
		{
			Path:  "about",
			Value: profile.About,
		},
	}

	_, err := docRef.Update(ctx, updates)
	if err != nil {
		return nil, errors.New("failed to update document :" + err.Error())
	}

	profile.ID = docRef.ID

	return profile, nil

}

func (repository *ProfileRepositoryImpl) Delete(ctx context.Context, client *firestore.Client, profileId string) error {

	docRef := client.Collection("profiles").Doc(profileId)

	_, err := docRef.Delete(ctx)
	if err != nil {
		return errors.New("failed to delete document :" + err.Error())
	}

	return nil
}
