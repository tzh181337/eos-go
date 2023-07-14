package system

import (
	eos "github.com/armoniax/eos-go"
)

// NewSetalimits sets the account limits. Requires signature from `amax@active` account.
func NewSetalimits(account eos.AccountName, ramBytes, netWeight, cpuWeight int64) *eos.Action {
	a := &eos.Action{
		Account: AN("amax"),
		Name:    ActN("setalimit"),
		Authorization: []eos.PermissionLevel{
			{Actor: eos.AccountName("amax"), Permission: PN("active")},
		},
		ActionData: eos.NewActionData(Setalimits{
			Account:   account,
			RAMBytes:  ramBytes,
			NetWeight: netWeight,
			CPUWeight: cpuWeight,
		}),
	}
	return a
}

// Setalimits represents the `amax.system::setalimit` action.
type Setalimits struct {
	Account   eos.AccountName `json:"account"`
	RAMBytes  int64           `json:"ram_bytes"`
	NetWeight int64           `json:"net_weight"`
	CPUWeight int64           `json:"cpu_weight"`
}
