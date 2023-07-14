package system

import (
	"github.com/armoniax/eos-go"
)

func init() {
	eos.RegisterAction(AN("amax"), ActN("setcode"), SetCode{})
	eos.RegisterAction(AN("amax"), ActN("setabi"), SetABI{})
	eos.RegisterAction(AN("amax"), ActN("newaccount"), NewAccount{})
	eos.RegisterAction(AN("amax"), ActN("delegatebw"), DelegateBW{})
	eos.RegisterAction(AN("amax"), ActN("undelegatebw"), UndelegateBW{})
	eos.RegisterAction(AN("amax"), ActN("refund"), Refund{})
	eos.RegisterAction(AN("amax"), ActN("regproducer"), RegProducer{})
	eos.RegisterAction(AN("amax"), ActN("unregprod"), UnregProducer{})
	eos.RegisterAction(AN("amax"), ActN("regproxy"), RegProxy{})
	eos.RegisterAction(AN("amax"), ActN("voteproducer"), VoteProducer{})
	eos.RegisterAction(AN("amax"), ActN("claimrewards"), ClaimRewards{})
	eos.RegisterAction(AN("amax"), ActN("buyram"), BuyRAM{})
	eos.RegisterAction(AN("amax"), ActN("buyrambytes"), BuyRAMBytes{})
	eos.RegisterAction(AN("amax"), ActN("linkauth"), LinkAuth{})
	eos.RegisterAction(AN("amax"), ActN("unlinkauth"), UnlinkAuth{})
	eos.RegisterAction(AN("amax"), ActN("deleteauth"), DeleteAuth{})
	eos.RegisterAction(AN("amax"), ActN("rmvproducer"), RemoveProducer{})
	eos.RegisterAction(AN("amax"), ActN("setprods"), SetProds{})
	eos.RegisterAction(AN("amax"), ActN("setpriv"), SetPriv{})
	eos.RegisterAction(AN("amax"), ActN("canceldelay"), CancelDelay{})
	eos.RegisterAction(AN("amax"), ActN("bidname"), Bidname{})
	// eos.RegisterAction(AN("amax"), ActN("nonce"), &Nonce{})
	eos.RegisterAction(AN("amax"), ActN("sellram"), SellRAM{})
	eos.RegisterAction(AN("amax"), ActN("updateauth"), UpdateAuth{})
	eos.RegisterAction(AN("amax"), ActN("setramrate"), SetRAMRate{})
	eos.RegisterAction(AN("amax"), ActN("setalimits"), Setalimits{})
}

var AN = eos.AN
var PN = eos.PN
var ActN = eos.ActN
