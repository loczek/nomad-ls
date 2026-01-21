package languages

type LanguageID string

const (
	NomadJob LanguageID = "nomad"
)

func (l LanguageID) String() string {
	return string(l)
}
