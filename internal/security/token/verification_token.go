package token

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type VerificationToken struct {
	Token     string
	ExpiresIn time.Time
}

func GenerateVerificationCode() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
