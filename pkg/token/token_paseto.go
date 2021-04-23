package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

// PasetoToken is a PASETO token
type PasetoToken struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

// NewPasetoToken creates a new PasetoToken
func NewPaseto(symmetricKey string) (Token, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoToken{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

// CreateToken creates a new token for a specific username and duration
func (p *PasetoToken) CreateToken(username string, scope string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, scope, duration)
	if err != nil {
		return "", err
	}

	return p.paseto.Encrypt(p.symmetricKey, payload, nil)
}

// VerifyToken checks if the token is valid or not
func (p *PasetoToken) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := p.paseto.Decrypt(token, p.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
