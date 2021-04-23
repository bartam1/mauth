package domain

import "time"

type User struct {
	UserName    string
	UserSecret  string
	FullName    string
	Scopes      []string
	CreatedAt   time.Time
	LastLogTime time.Time
	LastLogIp   string
}
