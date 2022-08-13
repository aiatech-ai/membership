package token

import "time"

type SecurityToken struct {
	Token   string
	Payload *Payload
}

type AuthenticationToken interface {
	IssueToken(username string, duration time.Duration) (*SecurityToken, error)
	VerifyToken(token string) (*Payload, error)
}
