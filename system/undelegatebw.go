package system

import (
	eos "github.com/armoniax/eos-go"
)

// NewUndelegateBW returns a `undelegatebw` action that lives on the
// `amax.system` contract.
func NewUndelegateBW(from, receiver eos.AccountName, unstakeCPU, unstakeNet eos.Asset) *eos.Action {
	return &eos.Action{
		Account: AN("amax"),
		Name:    ActN("undelegatebw"),
		Authorization: []eos.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(UndelegateBW{
			From:       from,
			Receiver:   receiver,
			UnstakeNet: unstakeNet,
			UnstakeCPU: unstakeCPU,
		}),
	}
}

// UndelegateBW represents the `amax.system::undelegatebw` action.
type UndelegateBW struct {
	From       eos.AccountName `json:"from"`
	Receiver   eos.AccountName `json:"receiver"`
	UnstakeNet eos.Asset       `json:"unstake_net_quantity"`
	UnstakeCPU eos.Asset       `json:"unstake_cpu_quantity"`
}
