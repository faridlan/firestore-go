package educationservice

import (
	"context"

	"github.com/faridlan/firestore-go/model/web"
)

type EducationService interface {
	Save(ctx context.Context, request *web.EducationCreate) (*web.EducationResponse, error)
	Update(ctx context.Context, request *web.EducationCreate) (*web.EducationResponse, error)
	Find(ctx context.Context) ([]*web.EducationResponse, error)
	FindId(ctx context.Context, educationId string) (*web.EducationResponse, error)
	Delete(ctx context.Context, educationId string) error
}
