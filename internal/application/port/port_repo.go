package port

import (
	"context"

	"github.com/bartam1/mauth/internal/application/domain"
)

type RepoIf interface {
	UserAuth(ctx context.Context, ua domain.UserAuth) (err error)
}
