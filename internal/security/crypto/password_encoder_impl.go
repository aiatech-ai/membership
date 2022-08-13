package crypto

import "golang.org/x/crypto/bcrypt"

const defaultCost = bcrypt.DefaultCost

type passwordEncoder struct {
}

var _ PasswordEncoder = (*passwordEncoder)(nil)

func NewPasswordEncoder() *passwordEncoder {
	return &passwordEncoder{}
}

func (p *passwordEncoder) Encode(rawPassword string) (string, error) {
	encodedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), defaultCost)
	return string(encodedPassword), err
}

func (p *passwordEncoder) Matches(rawPassword, encodedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encodedPassword), []byte(rawPassword)) == nil
}
