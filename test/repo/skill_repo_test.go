package repo

import (
	"context"
	"fmt"
	"testing"

	"github.com/faridlan/firestore-go/config"
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
