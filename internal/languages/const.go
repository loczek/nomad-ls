package languages

type LanguageID string

const (
	NomadJob           LanguageID = "nomad-job"
	NomadAgent         LanguageID = "nomad-agent"
	NomadACL           LanguageID = "nomad-acl"
	NomadNapespace     LanguageID = "nomad-namespace"
	NomadNodePool      LanguageID = "nomad-node-pool"
	NomadResourceQuota LanguageID = "nomad-resource-quota"
	NomadVariable      LanguageID = "nomad-variable"
	NomadVolume        LanguageID = "nomad-volume"
)

func (l LanguageID) String() string {
	return string(l)
}
