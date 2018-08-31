package dbv2

import (
	"github.com/iost-official/Go-IOS-Protocol/common"
	"github.com/iost-official/explorer/backend/model/blkchain"
)

type ActionRaw struct {
	Contract   string `bson:"contract"`
	ActionName string `bson:"actionName"`
	Data       string `bson:"data"`
}

type SignatureRaw struct {
	Algorithm int32  `bson:"algorithm"`
	Sig       string `bson:"sig"`
	PubKey    string `bson:"pubKey"`
}

type Tx struct {
	BlockNumber     int64           `bson:"blockNumber"`
	Time       int64           `bson:"time"`
	Hash       string          `bson:"hash"`
	Expiration int64           `bson:"expiration"`
	GasPrice   int64           `bson:"gasPrice"`
	GasLimit   int64           `bson:"gasLimit"`
	Actions    []ActionRaw    `bson:"actions"`
	Signers    []string        `bson:"signers"`
	Signs      []SignatureRaw `bson:"signs"`
	Publisher  SignatureRaw   `bson:"publisher"`
}

func RpcGetTxByHash(txHash string) (*Tx, error) {
	txRes, err := blkchain.GetTxByHash(txHash)
	if err != nil {
		return nil, err
	}
	txRaw := txRes.TxRaw
	actions := make([]ActionRaw, len(txRaw.Actions))
	for i, v := range txRaw.Actions {
		actions[i] = ActionRaw{
			Contract:   v.Contract,
			ActionName: v.ActionName,
			Data:       v.Data,
		}
	}
	publisher := SignatureRaw{
		Algorithm: txRaw.Publisher.Algorithm,
		Sig:       common.Base58Encode(txRaw.Publisher.Sig),
		PubKey:    common.Base58Encode(txRaw.Publisher.PubKey),
	}
	signs := make([]SignatureRaw, len(txRaw.Signs))
	for i, v := range txRaw.Signs {
		signs[i] = SignatureRaw{
			Algorithm: v.Algorithm,
			Sig:       common.Base58Encode(v.Sig),
			PubKey:    common.Base58Encode(v.PubKey),
		}
	}
	return &Tx{
		Time:       txRaw.Time,
		Hash:       txHash,
		Expiration: txRaw.Expiration,
		GasPrice:   txRaw.GasPrice,
		GasLimit:   txRaw.GasLimit,
		Actions:    actions,
		Signers:    byteSliceArrayToStringArray(txRaw.Signers),
		Signs:      signs,
		Publisher:  publisher,
	}, nil
}

func byteSliceArrayToStringArray(origin [][]byte) []string {
	vsm := make([]string, len(origin))
	for i, v := range origin {
		vsm[i] = common.Base58Encode(v)
	}
	return vsm
}
