package cron

import (
	"github.com/globalsign/mgo/bson"
	"github.com/iost-official/explorer/backend/model/db"
	"github.com/iost-official/explorer/backend/model/dbv2"
	"log"
	"sync"
	"time"
)

func UpdateTxns(wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(time.Second * 2)
	var start int = 0

	txnC, err := db.GetCollection(dbv2.CollectionTxs)

	if err != nil {
		log.Println("UpdateTxns get txns collection error:", err)
		return
	}

	flatxnC, err := db.GetCollection(dbv2.CollectionFlatTx)

	if err != nil {
		log.Println("UpdateTxns get flatxs collection error:", err)
	}

	for range ticker.C {
		step := 300
		var txns = make([]*dbv2.Tx, 0)

		err = txnC.Find(bson.M{}).Sort("_id").Skip(start).Limit(step).All(&txns)

		if err != nil {
			log.Println("UpdateTxns query txns collection error:", err)
			return
		}

		for _, txn := range txns {
			originTxn := *txn
			newTxn, err := dbv2.RpcGetTxByHash(originTxn.Hash)

			if err != nil {
				log.Println("UpdateTxns RpcGetTxByHash error:", err)
				continue
			}

			flatxs := newTxn.ToFlatTx()
			var tmpFlatx dbv2.FlatTx

			for _, tx := range flatxs {
				err := flatxnC.Find(bson.M{"actionIndex": tx.ActionIndex,
					"hash": tx.Hash}).One(&tmpFlatx)

				if err != nil {
					flatxnC.Insert(tx)
				} else {
					continue
				}
			}

			if originTxn.Time != 0 {
				continue
			}

			txnC.Update(bson.M{"hash": newTxn.Hash},
				bson.M{
					"$set": bson.M{
						"time":       newTxn.Time,
						"expiration": newTxn.Expiration,
						"gasPrice":   newTxn.GasPrice,
						"gasLimit":   newTxn.GasLimit,
						"actions":    newTxn.Actions,
						"signers":    newTxn.Signers,
						"signs":      newTxn.Signs,
						"publisher":  newTxn.Publisher}})
		}

		start += step
		log.Println("start = ", start)
	}
}
