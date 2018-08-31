package dbv2

import (
	"encoding/json"
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
	BlockNumber int64          `bson:"blockNumber"`
	Time        int64          `bson:"time"`
	Hash        string         `bson:"hash"`
	Expiration  int64          `bson:"expiration"`
	GasPrice    int64          `bson:"gasPrice"`
	GasLimit    int64          `bson:"gasLimit"`
	Actions     []ActionRaw    `bson:"actions"`
	Signers     []string       `bson:"signers"`
	Signs       []SignatureRaw `bson:"signs"`
	Publisher   SignatureRaw   `bson:"publisher"`
}

// 将 Tx.Actions 打平后的数据结构， 如果actionName == Transfer 则会解析出 from, to, amount
type FlatTx struct {
	BlockNumber int64          `bson:"blockNumber"`
	Time        int64          `bson:"time"`
	Hash        string         `bson:"hash"`
	Expiration  int64          `bson:"expiration"`
	GasPrice    int64          `bson:"gasPrice"`
	GasLimit    int64          `bson:"gasLimit"`
	Action      ActionRaw      `bson:"action"`
	Signers     []string       `bson:"signers"`
	Signs       []SignatureRaw `bson:"signs"`
	Publisher   SignatureRaw   `bson:"publisher"`
	From        string         `bson:"from"`
	To          string         `bson:"to"`
	Amount      float64          `bson:"amount"`  // 转发数量
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


func (tx *Tx) ToFlatTx() []FlatTx {
	flatTx := make([]FlatTx, len(tx.Actions))

	for i, v := range tx.Actions {
		var from, to string
		var amount float64
		if v.ActionName == "Transfer" {
			var tmp []interface{}
			json.Unmarshal([]byte(v.Data), &tmp)  // TODO check error
			from = tmp[0].(string)
			to = tmp[1].(string)
			amount = tmp[2].(float64)
		}
		flatTx[i] = FlatTx{
			BlockNumber: tx.BlockNumber,
			Time: tx.Time,
			Hash: tx.Hash,
			Expiration: tx.Expiration,
			GasPrice: tx.GasPrice,
			GasLimit: tx.GasLimit,
			Signers: tx.Signers,
			Publisher: tx.Publisher,
			Signs: tx.Signs,
			Action: v,
			From: from,
			To: to,
			Amount: amount,
		}
	}
	return flatTx
}


func byteSliceArrayToStringArray(origin [][]byte) []string {
	vsm := make([]string, len(origin))
	for i, v := range origin {
		vsm[i] = common.Base58Encode(v)
	}
	return vsm
}
