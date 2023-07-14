package system

import eos "github.com/armoniax/eos-go"

// NewSetPriv returns a `setpriv` action that lives on the
// `amax.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `eos-bios` boot process by the
// `amax.system` contract.
func NewSetPriv(account eos.AccountName) *eos.Action {
	a := &eos.Action{
		Account: AN("amax"),
		Name:    ActN("setpriv"),
		Authorization: []eos.PermissionLevel{
			{Actor: AN("amax"), Permission: PN("active")},
		},
		ActionData: eos.NewActionData(SetPriv{
			Account: account,
			IsPriv:  eos.Bool(true),
		}),
	}
	return a
}

// SetPriv sets privileged account status. Used in the bios boot mechanism.
type SetPriv struct {
	Account eos.AccountName `json:"account"`
	IsPriv  eos.Bool        `json:"is_priv"`
}
