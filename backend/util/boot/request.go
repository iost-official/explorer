//
// request.go
// Copyright (C) 2018 jack <jack@iOSTdeMacBook-Pro.local>
//
// Distributed under terms of the MIT license.
//

package boot

import (
	"errors"
	"strings"

	"github.com/iost-official/go-iost/account"
	"github.com/iost-official/go-iost/common"
	"github.com/iost-official/go-iost/crypto"
	"github.com/iost-official/go-iost/iwallet"
)

type Request struct {
	Net    string
	PubKey string
	Name   string
	MK     string
}

func f(r rune) bool {
	return (r < '0') || ((r > '9') && (r < 'A')) || ((r > 'Z') && (r < 'a')) || (r > 'z')
}

func (r *Request) IsValid() error {

	if len(r.Net) < 1 {
		return errors.New("no net name")
	}
	found := false
	for k := range C.Net {
		if k == r.Net {
			found = true
			break
		}
	}
	if !found {
		return errors.New("net not found")
	}

	if len(r.Name) < 5 || len(r.Name) > 11 {
		return errors.New("invalid account string")
	}

	if strings.IndexFunc(r.PubKey, f) != -1 {
		return errors.New("invalid pubkey")
	}

	return nil
}

func (r *Request) CreateAccount() (string, error) {
	acc, err := account.NewKeyPair(common.Base58Decode(C.Net[r.Net].AdminKey), crypto.Ed25519)
	if err != nil {
		return "", err
	}

	var sdk = iwallet.SDK{}
	sdk.SetServer(C.Net[r.Net].EndPoint)
	sdk.SetAccount(C.Net[r.Net].AdminAccount, acc)
	sdk.SetTxInfo(1000000, 1, 60, 0)
	sdk.SetCheckResult(true, 3, 10)
	sdk.SetAmountLimit("*:unlimited")
	sdk.SetChainID(C.Net[r.Net].ChainID)

	pubKey := r.PubKey
	txHash, err := sdk.CreateNewAccount(r.Name, pubKey, pubKey,
		C.Net[r.Net].Initial.GasPledge, C.Net[r.Net].Initial.Ram, C.Net[r.Net].Initial.Coin)
	if err != nil {
		return "", err
	}

	return txHash, nil
}
