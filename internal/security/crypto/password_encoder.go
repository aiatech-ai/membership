package crypto

type PasswordEncoder interface {
	Encode(rawPassword string) (string, error)
	Matches(rawPassword, encodedPassword string) bool
}
