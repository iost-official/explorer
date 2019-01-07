package db

import "github.com/iost-official/explorer/backend/model/blockchain/rpcpb"

type ContractTx struct {
	ID     string `bson:"id"`
	Time   int64  `bson:"time"`
	TxHash string `bson:"txHash"`
}

type Contract struct {
	ID           string          `bson:"id"`
	Domain       string          `bson:"domain"`
	CreateTime   int64           `bson:"createTime"`
	Creator      string          `bson:"creator"`
	Balance      float64         `bson:"balance"`
	ContractInfo *rpcpb.Contract `bson:"contractInfo"`
}

func NewContract(id string, time int64, creator string) *Contract {
	return &Contract{
		ID:         id,
		CreateTime: time,
		Creator:    creator,
	}
}
