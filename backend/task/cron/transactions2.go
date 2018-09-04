package cron

import (
	"github.com/globalsign/mgo/bson"
	"github.com/iost-official/explorer/backend/model2/db"
	"log"
	"sync"
	"time"
)

func UpdateTxns(wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(time.Second * 2)

	txnC, err := db.GetCollection(db.CollectionTxs)

	if err != nil {
		log.Println("UpdateTxns get txns collection error:", err)
		return
	}

	flatxnC, err := db.GetCollection(db.CollectionFlatTx)

	if err != nil {
		log.Println("UpdateTxns get flatxs collection error:", err)
		return
	}

	for range ticker.C {
		step := 300
		var txns = make([]*db.Tx, 0)

		err = txnC.Find(bson.M{"time": bson.M{"$eq": 0}}).
			Sort("_id").Limit(step).All(&txns)

		if err != nil {
			log.Println("UpdateTxns query txns collection error:", err)
			return
		}

		for _, txn := range txns {
			newTxn, err := db.RpcGetTxByHash(txn.Hash)

			if err != nil {
				log.Println("UpdateTxns RpcGetTxByHash error:", err)
				continue
			}

			flatxns := newTxn.ToFlatTx()
			var tmpFlatxn db.FlatTx

			for _, tx := range flatxns {
				errFind := flatxnC.Find(bson.M{"actionIndex": tx.ActionIndex,
					"hash": tx.Hash}).One(&tmpFlatxn)

				// flatxn not found
				if errFind != nil {
					errInsert := flatxnC.Insert(*tx)

					if errInsert != nil {
						log.Println("failed to insert to flatxnC")
						//	TODO: save those record to database to try again
					}
				}
			}

			errUpdate := txnC.Update(bson.M{"hash": newTxn.Hash},
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

			if errUpdate != nil {
				log.Println("failed to update txn")
				//	TODO: save failed record to database to try again
			}
		}
	}
}
