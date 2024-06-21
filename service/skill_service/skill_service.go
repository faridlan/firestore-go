package skillservice

import (
	"context"

	"github.com/faridlan/firestore-go/model/web"
)

type SkillService interface {
	Save(ctx context.Context, request *web.SkillCreate) (*web.SkillResponse, error)
	Update(ctx context.Context, request *web.SkillCreate) (*web.SkillResponse, error)
	Find(ctx context.Context) ([]*web.SkillResponse, error)
	FindId(ctx context.Context, SkillId string) (*web.SkillResponse, error)
	Delete(ctx context.Context, SkillId string) error
}
