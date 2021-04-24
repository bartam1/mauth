package domain

const (
	AuthFailed = iota
	AuthSuccess
)

const (
	UTypePerson = iota
	UTypeService
)

const (
	ScopeRead = iota
	ScopeWrite
	ScopeDelete
)
