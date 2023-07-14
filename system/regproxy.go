package system

import (
	eos "github.com/armoniax/eos-go"
)

// NewRegProxy returns a `regproxy` action that lives on the
// `amax.system` contract.
func NewRegProxy(proxy eos.AccountName, isProxy bool) *eos.Action {
	return &eos.Action{
		Account: AN("amax"),
		Name:    ActN("regproxy"),
		Authorization: []eos.PermissionLevel{
			{Actor: proxy, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(RegProxy{
			Proxy:   proxy,
			IsProxy: isProxy,
		}),
	}
}

// RegProxy represents the `amax.system::regproxy` action
type RegProxy struct {
	Proxy   eos.AccountName `json:"proxy"`
	IsProxy bool            `json:"isproxy"`
}
