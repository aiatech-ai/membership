package enum

type UserUpdateReason int

const (
	CREATE_NEW UserUpdateReason = iota + 1
	UPDATE_FULL_NAME
	UPDATE_PASSWORD
	UPDATE_OTP
	UPDATE_TO_ACTIVE
	UPDATE_TO_DEACTIVE
)

func (s UserUpdateReason) String() string {
	return [...]string{
		"CREATE_NEW",
		"UPDATE_FULL_NAME",
		"UPDATE_PASSWORD",
		"UPDATE_OTP",
		"UPDATE_TO_ACTIVE",
		"UPDATE_TO_DEACTIVE",
	}[s]
}
