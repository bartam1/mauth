package service

import (
	"context"

	"github.com/bartam1/mauth/internal/application/domain"
	"github.com/bartam1/mauth/internal/application/port"
)

type Entity struct {
	db port.RepoIf
}

func New(db port.RepoIf) Entity {
	return Entity{db}
}
func (e *Entity) UserAuth(ctx context.Context, ua domain.UserAuth) error {
	return nil
}
func (e *Entity) LoginPage() (s string, err error) {
	return "", nil
}
