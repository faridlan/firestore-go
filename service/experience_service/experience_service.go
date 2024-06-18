package experienceservice

import (
	"context"

	"github.com/faridlan/firestore-go/model/web"
)

type ExperienceService interface {
	Save(ctx context.Context, request *web.ExperienceCreate) (*web.ExperienceResponse, error)
	Update(ctx context.Context, request *web.ExperienceCreate) (*web.ExperienceResponse, error)
	FindId(ctx context.Context, experienceId string) (*web.ExperienceResponse, error)
	Find(ctx context.Context) ([]*web.ExperienceResponse, error)
	Delete(ctx context.Context, experienceId string) error
}
