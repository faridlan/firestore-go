package skillrepository

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/faridlan/firestore-go/model/domain"
	"google.golang.org/api/iterator"
)

type SkillRepositoryImpl struct {
}

func NewSkillRepository() SkillRepository {
	return &SkillRepositoryImpl{}
}

func (repository *SkillRepositoryImpl) Save(ctx context.Context, client *firestore.Client, skill *domain.Skill) (*domain.Skill, error) {

	docRef, _, err := client.Collection("skills").Add(ctx, skill)
	if err != nil {
		return nil, errors.New("failed to create new doc : " + err.Error())
	}

	skill.ID = docRef.ID

	return skill, err

}

func (repository *SkillRepositoryImpl) Update(ctx context.Context, client *firestore.Client, skill *domain.Skill) (*domain.Skill, error) {

	docRef := client.Collection("skills").Doc(skill.ID)

	updates := []firestore.Update{
		{
			Path:  "name",
			Value: skill.Name,
		},
	}

	_, err := docRef.Update(ctx, updates)
	if err != nil {
		return nil, errors.New("failed to update doc : " + err.Error())
	}

	skill.ID = docRef.ID

	return skill, nil

}

func (repository *SkillRepositoryImpl) Find(ctx context.Context, client *firestore.Client) ([]*domain.Skill, error) {

	iter := client.Collection("skills").Documents(ctx)

	defer iter.Stop()

	skills := []*domain.Skill{}

	for {

		docSnapshot, err := iter.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, errors.New("failed to iterate doc : " + err.Error())
		}

		skill := &domain.Skill{}
		err = docSnapshot.DataTo(&skill)
		if err != nil {
			return nil, errors.New("failed to decode doc : " + err.Error())
		}

		skill.ID = docSnapshot.Ref.ID

		skills = append(skills, skill)

	}

	return skills, nil

}

func (repository *SkillRepositoryImpl) FindId(ctx context.Context, client *firestore.Client, skillId string) (*domain.Skill, error) {

	docRef := client.Collection("skills").Doc(skillId)

	docSnapshot, err := docRef.Get(ctx)
	if err != nil {
		return nil, errors.New("skill not found")
	}

	skill := &domain.Skill{}
	err = docSnapshot.DataTo(&skill)
	if err != nil {
		return nil, errors.New("failed to decode doc : " + err.Error())
	}

	skill.ID = docSnapshot.Ref.ID

	return skill, nil

}

func (repository *SkillRepositoryImpl) Delete(ctx context.Context, client *firestore.Client, skillId string) error {

	docRef := client.Collection("skills").Doc(skillId)

	_, err := docRef.Delete(ctx)
	if err != nil {
		return errors.New("failed to delete doc : " + err.Error())
	}

	return nil

}
