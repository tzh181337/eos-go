package system

import "github.com/armoniax/eos-go"

// NewNonce returns a `nonce` action that lives on the
// `amax.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `eos-bios` boot process by the
// `amax.system` contract.
func NewVoteProducer(voter eos.AccountName, proxy eos.AccountName, producers ...eos.AccountName) *eos.Action {
	a := &eos.Action{
		Account: AN("amax"),
		Name:    ActN("voteproducer"),
		Authorization: []eos.PermissionLevel{
			{Actor: voter, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(
			VoteProducer{
				Voter:     voter,
				Proxy:     proxy,
				Producers: producers,
			},
		),
	}
	return a
}

// VoteProducer represents the `amax.system::voteproducer` action
type VoteProducer struct {
	Voter     eos.AccountName   `json:"voter"`
	Proxy     eos.AccountName   `json:"proxy"`
	Producers []eos.AccountName `json:"producers"`
}
