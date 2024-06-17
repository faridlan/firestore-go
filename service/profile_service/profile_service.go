package profileservice

import (
	"context"

	"github.com/faridlan/firestore-go/model/web"
)

type ProfileService interface {
	Save(ctx context.Context, request *web.Profile) (*web.Profile, error)
	Find(ctx context.Context, profileId string) (*web.Profile, error)
	Update(ctx context.Context, request *web.Profile) (*web.Profile, error)
	Delete(ctx context.Context, profileId string) error
}
