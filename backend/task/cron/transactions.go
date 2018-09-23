package cron

import (
	"github.com/globalsign/mgo/bson"
	"github.com/iost-official/explorer/backend/model/db"
	"log"
	"sync"
	"time"
)

func UpdateTxns(wg *sync.WaitGroup, mark int) {
	defer wg.Done()

	ticker := time.NewTicker(time.Second)

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

	// first block contains only ONE transaction
	tx1stSynced := false

	for range ticker.C {
		step := 2000
		var txns = make([]*db.Tx, 0)

		query := bson.M{"time": bson.M{"$eq": 0}, "mark": mark}

		if tx1stSynced {
			// skipt block 0
			query = bson.M{
				"blockNumber": bson.M{"$ne": 0},
				"time":        bson.M{"$eq": 0},
				"mark":        mark}
		}

		err = txnC.Find(query).
			Sort("_id").Limit(step).All(&txns)

		if err != nil {
			log.Println("UpdateTxns query txns collection error:", err)
			continue
		}

		var flatxs []interface{}

		for _, txn := range txns {
			newTxn, err := db.RpcGetTxByHash(txn.Hash)

			if err != nil {
				log.Println("UpdateTxns RpcGetTxByHash error:", err)
				continue
			}

			newTxn.BlockNumber = txn.BlockNumber

			flatxns := newTxn.ToFlatTx()
			var tmpFlatxn db.FlatTx

			for _, tx := range flatxns {
				errFind := flatxnC.Find(bson.M{"actionIndex": tx.ActionIndex,
					"hash": tx.Hash}).One(&tmpFlatxn)

				// flatxn not found
				if errFind != nil {
					flatxs = append(flatxs, *tx)
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
				continue
			}

			if newTxn.BlockNumber == 0 {
				tx1stSynced = true
			}
		}
		if len(flatxs) != 0 {
			err := flatxnC.Insert(flatxs...)
			if nil != err{
				log.Println("fail to insert flatxs, err: ", err)
			} else {
				log.Println("update flatxs, size: ", len(flatxs))
			}
		}
		log.Println("update txs, size: ", len(txns))
	}
}
