package system

import "github.com/armoniax/eos-go"

// NewNonce returns a `nonce` action that lives on the
// `amax.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `eos-bios` boot process by the
// `amax.system` contract.
func NewNonce(nonce string) *eos.Action {
	a := &eos.Action{
		Account:       AN("amax"),
		Name:          ActN("nonce"),
		Authorization: []eos.PermissionLevel{
			//{Actor: AN("amax"), Permission: PN("active")},
		},
		ActionData: eos.NewActionData(Nonce{
			Value: nonce,
		}),
	}
	return a
}
