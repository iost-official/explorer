package db

import (
	"github.com/globalsign/mgo/bson"
	"github.com/GincoInc/iost-explorer/backend/model/blockchain/rpcpb"
)

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

func GetContractTxByID(ID string, start, limit int) ([]*ContractTx, error) {
	contractTxC := GetCollection(CollectionContractTx)
	query := bson.M{
		"id": ID,
	}

	var contractTxList []*ContractTx
	err := contractTxC.Find(query).Sort("-time").Skip(start).Limit(limit).All(&contractTxList)
	if err != nil {
		return nil, err
	}
	return contractTxList, nil
}

func GetContractTxNumber(ID string) (int, error) {
	contractTxC := GetCollection(CollectionContractTx)
	query := bson.M{
		"id": ID,
	}

	return contractTxC.Find(query).Count()
}
