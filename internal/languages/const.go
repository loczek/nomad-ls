package languages

type LanguageID string

const (
	NomadACL               LanguageID = "nomad-acl"
	NomadAgent             LanguageID = "nomad-agent"
	NomadCSIVolume         LanguageID = "nomad-csi-volume"
	NomadDynamicHostVolume LanguageID = "nomad-dynamic-host-volume"
	NomadJob               LanguageID = "nomad-job"
	NomadNapespace         LanguageID = "nomad-namespace"
	NomadNodePool          LanguageID = "nomad-node-pool"
	NomadResourceQuota     LanguageID = "nomad-resource-quota"
	NomadVariable          LanguageID = "nomad-variable"
)

var langs = map[string]LanguageID{
	"nomad-acl":                 NomadACL,
	"nomad-agent":               NomadAgent,
	"nomad-csi-volume":          NomadCSIVolume,
	"nomad-dynamic-host-volume": NomadDynamicHostVolume,
	"nomad-job":                 NomadJob,
	"nomad-namespace":           NomadNapespace,
	"nomad-node-pool":           NomadNodePool,
	"nomad-resource-quota":      NomadResourceQuota,
	"nomad-variable":            NomadVariable,
}

func (l LanguageID) String() string {
	return string(l)
}

func NewFromString(id string) (LanguageID, error) {
	if val, ok := langs[id]; ok {
		return val, nil
	}

	// TODO: remove this in the future
	return NomadJob, nil

	// return "", fmt.Errorf("LanguageID: \"%s\" is not a valid language id", id)
}
