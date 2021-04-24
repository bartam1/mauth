package domain

import (
	"time"
)

type User struct {
	UserId       uint
	UserType     UserType
	UserName     UserName
	UserSecret   UserSecret
	FullName     FullName
	Scopes       []Scope
	RefreshToken string
	AccessTokens []string
	CreatedAt    time.Time
	LastLogTime  time.Time
	LastLogIp    string
	FailLogCount int
	FailLogIp    string
}
