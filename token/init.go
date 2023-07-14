package token

import "github.com/armoniax/eos-go"

func init() {
	eos.RegisterAction(AN("eamax.token"), ActN("transfer"), Transfer{})
	eos.RegisterAction(AN("eamax.token"), ActN("issue"), Issue{})
	eos.RegisterAction(AN("eamax.token"), ActN("create"), Create{})
}
