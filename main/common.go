package main

import (
	"context"
	"github.com/armoniax/eos-go"
)

//export SubmitTransaction
func SubmitTransaction(rpcNode, privateKey string, eosActions []*eos.Action) string {
	ctx := context.Background()
	keyBag := &eos.KeyBag{}
	err := keyBag.ImportAmaxPrivateKey(ctx, privateKey)
	if err != nil {
		return err.Error()
	}

	api := eos.New(rpcNode)
	api.SetSigner(keyBag)

	txOpts := &eos.TxOptions{}
	err = txOpts.FillFromChain(ctx, api)
	if err != nil {
		return err.Error()
	}

	tx := eos.NewTransaction(eosActions, txOpts)
	_, packedTrx, err := api.SignTransaction(ctx, tx, txOpts.ChainID, eos.CompressionNone)
	if err != nil {
		return err.Error()
	}

	out, err := api.PushTransaction(ctx, packedTrx)
	if err != nil {
		return err.Error()
	}

	return out.TransactionID
}

//export GenAction
func GenAction(contractName, actionName, submitter, permission string, obj any) *eos.Action {

	action := genEosAction(eos.AccountName(contractName), actionName, eos.AccountName(submitter), permission)
	if obj != nil {
		action.ActionData = eos.NewActionData(obj)
	}
	return action
}

func genEosAction(account eos.AccountName, name string, actor eos.AccountName, permission string) *eos.Action {
	return &eos.Action{
		Account: account,
		Name:    eos.ActN(name),
		Authorization: []eos.PermissionLevel{
			{Actor: actor, Permission: eos.PN(permission)},
		},
		ActionData: eos.ActionData{
			Data: "",
		},
	}
}

func main() {

}