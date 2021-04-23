package service

import (
	"context"

	"github.com/bartam1/mauth/internal/application/domain"
)

type Interface interface {
	UserAuth(context.Context, domain.UserAuth) error
	LoginPage() (s string, err error)
}
