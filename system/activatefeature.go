package system

import (
	"github.com/armoniax/eos-go"
)

func NewActivateFeature(featureDigest eos.Checksum256) *eos.Action {
	return &eos.Action{
		Account: AN("amax"),
		Name:    ActN("activate"),
		Authorization: []eos.PermissionLevel{
			{Actor: AN("amax"), Permission: PN("active")},
		},
		ActionData: eos.NewActionData(Activate{
			FeatureDigest: featureDigest,
		}),
	}
}

// Activate represents a `activate` action on the `amax` contract.
type Activate struct {
	FeatureDigest eos.Checksum256 `json:"feature_digest"`
}
