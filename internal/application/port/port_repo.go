package port

import (
	"context"

	"github.com/bartam1/mauth/internal/application/domain"
)

type RepoIf interface {
	UserAuth(ctx context.Context, ua domain.UserAuth, updateFn func(u *domain.User, s domain.AuthStatus) (*domain.User, error)) (s *domain.User, err error)
}
