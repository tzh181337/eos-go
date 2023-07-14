package system

import (
	eos "github.com/armoniax/eos-go"
	"github.com/armoniax/eos-go/ecc"
)

// NewSetPriv returns a `setpriv` action that lives on the
// `amax.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `eos-bios` boot process by the
// `amax.system` contract.
func NewSetProds(producers []ProducerKey) *eos.Action {
	a := &eos.Action{
		Account: AN("amax"),
		Name:    ActN("setprods"),
		Authorization: []eos.PermissionLevel{
			{Actor: AN("amax"), Permission: PN("active")},
		},
		ActionData: eos.NewActionData(SetProds{
			Schedule: producers,
		}),
	}
	return a
}

// SetProds is present in `amax.bios` contract. Used only at boot time.
type SetProds struct {
	Schedule []ProducerKey `json:"schedule"`
}

type ProducerKey struct {
	ProducerName    eos.AccountName `json:"producer_name"`
	BlockSigningKey ecc.PublicKey   `json:"block_signing_key"`
}
