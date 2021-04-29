package domain

import (
	"time"
)

type User struct {
	UserId       uint
	UserType     *UserType
	UserName     *UserName
	UserSecret   *UserSecret
	Scopes       *Scopes
	RefreshToken *string
	AccessTokens []*string
	CreatedAt    time.Time
	LastLogTime  []time.Time
	LastLogIp    []*string
	FailLogCount int
	FailLogTime  []time.Time
	FailLogIp    []*string
}

func NewUser(UserType UType, UserName string, UserSecret string) (*User, error) {
	ut, err := NewUserType(UserType)
	if err != nil {
		return nil, err
	}
	un, err := NewUserName(UserName)
	if err != nil {
		return nil, err
	}
	us, err := NewUserSecret(UserSecret)
	if err != nil {
		return nil, err
	}
	ns, err := NewScope(UScopeRead)
	if err != nil {
		return nil, err
	}
	usc := NewScopes(ns)
	if err != nil {
		return nil, err
	}
	lt := make([]time.Time, 5)
	li := make([]*string, 5)
	ft := make([]time.Time, 5)
	fi := make([]*string, 5)
	at := make([]*string, 5)
	return &User{
		UserId:       0,
		UserType:     ut,
		UserName:     un,
		UserSecret:   us,
		Scopes:       usc,
		RefreshToken: nil,
		AccessTokens: at,
		CreatedAt:    time.Now(),
		LastLogTime:  lt,
		LastLogIp:    li,
		FailLogCount: 0,
		FailLogTime:  ft,
		FailLogIp:    fi,
	}, nil

}

type UserExtra struct {
	UserId   uint
	FullName *FullName
	//UserEmail	*UserEmail
}
