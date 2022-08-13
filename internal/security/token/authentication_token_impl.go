package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type authenticationToken struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

var _ AuthenticationToken = (*authenticationToken)(nil)

func NewAuthenticationToken(symmetricKey string) (AuthenticationToken, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &authenticationToken{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

func (maker *authenticationToken) IssueToken(username string, duration time.Duration) (*SecurityToken, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return nil, err
	}

	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	return &SecurityToken{
		Token:   token,
		Payload: payload,
	}, err
}

func (maker *authenticationToken) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
