package service

import (
	"context"
	"errors"
	"time"

	"github.com/bartam1/mauth/internal/application/domain"
	"github.com/bartam1/mauth/internal/application/port"
	"github.com/bartam1/mauth/pkg/config"
	"github.com/bartam1/mauth/pkg/token"
)

type Entity struct {
	db port.RepoIf
}

func New(db port.RepoIf) Entity {
	return Entity{db}
}
func (e *Entity) UserAuth(ctx context.Context, ua domain.UserAuth, ip string) (t string, err error) {
	user, err := e.db.UserAuth(ctx, ua, func(u *domain.User, s domain.AuthStatus) (e *domain.User, err error) {
		if s == domain.AuthSuccess {
			tok, err := token.NewPaseto(config.Global.TOKEN_SYMMETRIC_KEY)
			if err != nil {
				return nil, err
			}
			tstring, err := tok.CreateToken(u.UserName.Get(), u.Scopes.GetStrArray(), config.Global.ACCESS_TOKEN_DURATION)
			if err != nil {
				return nil, err
			}
			u.LastLogIp = ip
			u.AccessTokens = append(u.AccessTokens, tstring)
			u.LastLogTime = time.Now()
			u.FailLogCount = 0

			return u, nil
		}
		if s == domain.AuthWrongSecret {
			u.FailLogCount += 1
			u.FailLogIp = ip
			return u, nil

		}
		return nil, errors.New("Wrong AuthStatus!")
	})
	if err != nil {
		return "", err
	}
	return user.AccessTokens[len(user.AccessTokens)-1], nil
}
func (e *Entity) LoginPage() (s string, err error) {
	return "", nil
}
