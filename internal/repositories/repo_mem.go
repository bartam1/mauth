package repo

import (
	"context"
	"errors"

	"github.com/bartam1/mauth/internal/application/domain"
	"github.com/pkg/errors"
)

type Memory struct {
	db []*domain.User
}

func NewMem() (*Memory, error) {
	ret := Memory{db: make([]*domain.User, 5)}
	e, err := domain.NewUser(domain.UTypePerson, "probauser", "almafa")
	if err != nil {
		return nil, err
	}

	ret.db = append(ret.db, e)
	return &ret, nil
}

func (m *Memory) UserStore(ctx context.Context, u domain.User) error {
	return nil
}

func (m *Memory) UserAuth(ctx context.Context, ua domain.UserAuth, updateFn func(u *domain.User, s domain.AuthStatus) (*domain.User, error)) (t *domain.User, err error) {
	for _, u := range m.db {
		//We need different UserNames then we dont have to care about UserType, or primary key should be type+name
		if u.UserName.Get() == ua.UserName {
			if u.UserSecret.CompareString(ua.UserSecret) {
				au, err := updateFn(u, domain.AuthSuccess)
				if err != nil {
					return nil, errors.Wrap(err, "Error when user auth succ but failed to update user!")
				}
				m.UserStore(ctx, *au)
				return au, nil
			}
			au, err := updateFn(u, domain.AuthWrongSecret)
			if err != nil {
				return nil, errors.Wrap(err, "Error when user auth failed but failed to update user!")
			}
			return au, errors.New("Wrong secret!")
		}
	}
	return nil, errors.New("No such user!")
}
