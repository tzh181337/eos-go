package system

import (
	eos "github.com/armoniax/eos-go"
)

func NewSetRAM(maxRAMSize uint64) *eos.Action {
	a := &eos.Action{
		Account: AN("amax"),
		Name:    ActN("setram"),
		Authorization: []eos.PermissionLevel{
			{
				Actor:      AN("amax"),
				Permission: eos.PermissionName("active"),
			},
		},
		ActionData: eos.NewActionData(SetRAM{
			MaxRAMSize: eos.Uint64(maxRAMSize),
		}),
	}
	return a
}

// SetRAM represents the hard-coded `setram` action.
type SetRAM struct {
	MaxRAMSize eos.Uint64 `json:"max_ram_size"`
}
