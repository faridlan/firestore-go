package repo

import (
	"context"
	"fmt"
	"testing"

	"github.com/faridlan/firestore-go/config"
	"github.com/faridlan/firestore-go/model/domain"
	skillrepository "github.com/faridlan/firestore-go/repository/skill_repository"
	"github.com/stretchr/testify/assert"
)

var (
	client, _ = config.NewDatabase()
	skillRepo = skillrepository.NewSkillRepository()
	ctx       = context.Background()
)

func TestSkillFind(t *testing.T) {

	skillResponse, err := skillRepo.Find(ctx, client)
	assert.Nil(t, err)

	fmt.Println(skillResponse)
}

func TestSkillFindId(t *testing.T) {

	skillResponse, err := skillRepo.FindId(ctx, client, "3nBpOEmlcIJVFEbQNS00")
	assert.Nil(t, err)

	fmt.Println(skillResponse)
}

func TestSkillCreate(t *testing.T) {

	skill := &domain.Skill{
		Name: "skill-unit-tes",
	}

	skillResponse, err := skillRepo.Save(ctx, client, skill)
	assert.Nil(t, err)

	fmt.Println(skillResponse)
}

func TestSkillUpdate(t *testing.T) {

	err := skillRepo.Delete(ctx, client, "3nBpOEmlcIJVFEbQNS00")
	assert.Nil(t, err)

}
