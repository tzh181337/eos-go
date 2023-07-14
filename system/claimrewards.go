package system

import (
	eos "github.com/armoniax/eos-go"
)

// NewClaimRewards will buy at current market price a given number of
// bytes of RAM, and grant them to the `receiver` account.
func NewClaimRewards(owner eos.AccountName) *eos.Action {
	a := &eos.Action{
		Account: AN("amax"),
		Name:    ActN("claimrewards"),
		Authorization: []eos.PermissionLevel{
			{Actor: owner, Permission: eos.PermissionName("active")},
		},
		ActionData: eos.NewActionData(ClaimRewards{
			Owner: owner,
		}),
	}
	return a
}

// ClaimRewards represents the `amax.system::claimrewards` action.
type ClaimRewards struct {
	Owner eos.AccountName `json:"owner"`
}
