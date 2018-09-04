package db

import (
	"github.com/globalsign/mgo/bson"
	"log"
)

func GetTxnDetailByHash(txHash string) (*Tx, error) {
	txnDC, err := GetCollection(CollectionTxs)
	if err != nil {
		log.Println("UpdateTxns get txns collection error:", err)
		return nil, err
	}

	query := bson.M{
		"hash": txHash,
	}
	var txn *Tx
	err = txnDC.Find(query).One(&txn)

	return txn, err
}