package languages

type LanguageID string

const (
	NomadJob   LanguageID = "nomad-job"
	NomadAgent LanguageID = "nomad-agent"
	NomadVar   LanguageID = "nomad-var"
)

func (l LanguageID) String() string {
	return string(l)
}
