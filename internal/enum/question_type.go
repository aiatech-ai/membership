package enum

type QuestionType int

const (
	SINGLE_CHOISE QuestionType = iota + 1
	MULTIPLE_CHOISE
	TRUE_FALSE
)

func (s QuestionType) String() string {
	return [...]string{
		"SINGLE_CHOISE",
		"MULTIPLE_CHOISE",
		"TRUE_FALSE",
	}[s]
}
