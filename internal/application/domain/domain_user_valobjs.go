package domain

import (
	"errors"
	"regexp"

	"github.com/bartam1/mauth/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

type AuthStatus int
type UType int
type UScope int
type UserType struct {
	id UType
}

func (u *UserType) Get() int {
	return u.id
}

func NewUserType(t UType) (u *UserType, err error) {
	u = new(UserType)
	switch t {
	case UTypePerson:
		u.id = UTypePerson
	case UTypeService:
		u.id = UTypeService
	default:
		return nil, errors.New("There is no that kind of User type! (Person,Service")

	}
	return u, nil
}

type UserName struct {
	username string
}

func (u *UserName) Get() string {
	return u.username
}

func NewUserName(un string) (u *UserName, err error) {
	u = new(UserName)
	re := regexp.MustCompile(`^(?=.{4,20}$)(?![_.])(?!.*[_.]{2})[a-zA-Z0-9._]+(?<![_.])$`)
	if re.MatchString(un) {
		return nil, errors.New("UserName must be: 4-20 char long, and doesn't start with _,.  there is no __,_.,.. inside at the end. Allowed chars: a-z, A-Z, 0-9 .,_")

	}
	u.username = un
	return u, nil

}

type UserSecret struct {
	usersecret string
}

func (u *UserSecret) Get() string {
	return u.usersecret
}
func (u *UserSecret) CompareString(s string) bool {
	passsalt := s + config.Global.PASS_SALT
	hpass := bcrypt.GenerateFromPassword([]byte(passsalt), 12)
	if hpass == u.usersecret {
		return true
	}
	return false
}
func NewUserSecret(p string) (u *UserSecret, err error) {
	u = new(UserSecret)
	re := regexp.MustCompile(`^(?=.*[A-Za-z])(?=.*\d)(?=.*[@$!%*#?&])[A-Za-z\d@$!%*#?&]{8,}$`)
	if re.MatchString(p) {
		return nil, errors.New("Password must be minimum eight characters, at least one letter, one number and one special character!")
	}
	passsalt := p + config.Global.PASS_SALT
	hpass := bcrypt.GenerateFromPassword([]byte(passsalt), 12)
	u.usersecret = hpass
	return u, nil
}

type Scope struct {
	name string
}

func (u *Scope) Get() string {
	return u.name
}

func NewScope(i UScope) (s *Scope, err error) {
	s = new(Scope)
	switch i {
	case UScopeRead:
		s.name = "Read"
	case UScopeWrite:
		s.name = "Write"
	case UScopeDelete:
		s.name = "Delete"
	default:
		return nil, errors.New("There is no such scope! (Read,Write,Delete)")
	}
	return s, nil
}

type Scopes struct {
	scopes []*Scope
}

func (s *Scopes) GetStrArray() (as []string) {
	as = make([]string, len(s.scopes))
	for _, e := range s.scopes {
		as = append(as, e.name)
	}
	return as
}

func NewScopes(s ...*Scope) (sc *Scopes) {
	sc = new(Scopes)
	for _, e := range s {
		sc.scopes = append(sc.scopes, e)
	}
	return sc
}

type FullName struct {
	fullname string
}

func (u *FullName) Get() string {
	return u.fullname
}

func NewFullName(s string) (u *FullName, err error) {
	u = new(FullName)
	re := regexp.MustCompile(`^([\w]{3,})+\s+([\w\s]{3,})+$`)
	if re.MatchString(s) {
		return nil, errors.New("Not valid FullName!")

	}
	u.fullname = s
	return u, nil

}
