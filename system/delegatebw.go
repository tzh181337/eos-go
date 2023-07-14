package system

import (
	eos "github.com/armoniax/eos-go"
)

// NewDelegateBW returns a `delegatebw` action that lives on the
// `amax.system` contract.
func NewDelegateBW(from, receiver eos.AccountName, stakeCPU, stakeNet eos.Asset, transfer bool) *eos.Action {
	return &eos.Action{
		Account: AN("amax"),
		Name:    ActN("delegatebw"),
		Authorization: []eos.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(DelegateBW{
			From:     from,
			Receiver: receiver,
			StakeNet: stakeNet,
			StakeCPU: stakeCPU,
			Transfer: eos.Bool(transfer),
		}),
	}
}

// DelegateBW represents the `amax.system::delegatebw` action.
type DelegateBW struct {
	From     eos.AccountName `json:"from"`
	Receiver eos.AccountName `json:"receiver"`
	StakeNet eos.Asset       `json:"stake_net_quantity"`
	StakeCPU eos.Asset       `json:"stake_cpu_quantity"`
	Transfer eos.Bool        `json:"transfer"`
}
