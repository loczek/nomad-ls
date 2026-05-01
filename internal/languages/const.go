package languages

import (
	"fmt"
)

type LanguageID string

const (
	NomadACL           LanguageID = "nomad-acl"
	NomadAgent         LanguageID = "nomad-agent"
	NomadJob           LanguageID = "nomad-job"
	NomadNapespace     LanguageID = "nomad-namespace"
	NomadNodePool      LanguageID = "nomad-node-pool"
	NomadResourceQuota LanguageID = "nomad-resource-quota"
	NomadVariable      LanguageID = "nomad-variable"
	NomadVolumeCSI     LanguageID = "nomad-volume-csi"
	NomadVolumeDynamic LanguageID = "nomad-volume-dynamic"
)

var langs = map[string]LanguageID{
	"nomad-acl":            NomadACL,
	"nomad-agent":          NomadAgent,
	"nomad-job":            NomadJob,
	"nomad-namespace":      NomadNapespace,
	"nomad-node-pool":      NomadNodePool,
	"nomad-resource-quota": NomadResourceQuota,
	"nomad-variable":       NomadVariable,
	"nomad-volume-csi":     NomadVolumeCSI,
	"nomad-volume-dynamic": NomadVolumeDynamic,
}

func (l LanguageID) String() string {
	return string(l)
}

func NewFromString(id string) (LanguageID, error) {
	if val, ok := langs[id]; ok {
		return val, nil
	}

	return "", fmt.Errorf("LanguageID: \"%s\" is not a valid language id", id)
}
