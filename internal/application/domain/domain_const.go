package domain

const (
	AuthSuccess = iota
	AuthWrongUser
	AuthWrongSecret
)

const (
	UTypePerson = iota
	UTypeService
)

const (
	UScopeRead = iota
	UScopeWrite
	UScopeDelete
)
