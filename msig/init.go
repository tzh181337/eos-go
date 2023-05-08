package msig

import (
	"github.com/armoniax/eos-go"
)

func init() {
	eos.RegisterAction(AN("amax.msig"), ActN("propose"), &Propose{})
	eos.RegisterAction(AN("amax.msig"), ActN("approve"), &Approve{})
	eos.RegisterAction(AN("amax.msig"), ActN("unapprove"), &Unapprove{})
	eos.RegisterAction(AN("amax.msig"), ActN("cancel"), &Cancel{})
	eos.RegisterAction(AN("amax.msig"), ActN("exec"), &Exec{})
}

var AN = eos.AN
var PN = eos.PN
var ActN = eos.ActN
