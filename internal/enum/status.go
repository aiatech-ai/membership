package enum

type Status int

const (
	ACTIVED Status = iota + 1
	IN_ACTIVED
	DELETED
)

func (s Status) String() string {
	return [...]string{
		"ACTIVE",
		"IN_ACTIVE",
		"DELETE",
	}[s]
}
