package enum

type OrganizationType int

const (
	SCHOOL OrganizationType = iota + 1
	CENTER
	COMPANY
)

func (s OrganizationType) String() string {
	return [...]string{
		"SCHOOL",
		"CENTER",
		"COMPANY",
	}[s]
}
