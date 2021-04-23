package repo

import (
	"context"

	"github.com/bartam1/mauth/internal/application/domain"
)

type Memory struct {
	db map[string]string
}

func New() (Memory, error) {
	return Memory{db: map[string]string{}}, nil
}

func (m *Memory) UserAuth(ctx context.Context, ua domain.UserAuth) (err error) {
	return nil
}
