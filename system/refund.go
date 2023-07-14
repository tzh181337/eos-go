package system

import (
	eos "github.com/armoniax/eos-go"
)

// NewRefund returns a `refund` action that lives on the
// `amax.system` contract.
func NewRefund(owner eos.AccountName) *eos.Action {
	return &eos.Action{
		Account: AN("amax"),
		Name:    ActN("refund"),
		Authorization: []eos.PermissionLevel{
			{Actor: owner, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(Refund{
			Owner: owner,
		}),
	}
}

// Refund represents the `amax.system::refund` action
type Refund struct {
	Owner eos.AccountName `json:"owner"`
}
