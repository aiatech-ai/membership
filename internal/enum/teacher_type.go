package enum

type TeacherType int

const (
	MAIN TeacherType = iota + 1
	ASSISTANT
)

func (s TeacherType) String() string {
	return [...]string{
		"MAIN",
		"ASSISTANT",
	}[s]
}
