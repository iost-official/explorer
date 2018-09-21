package cron

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/iost-official/explorer/backend/model/db"
	"log"
	"sync"
	"time"
)

func UpdateTxns(wg *sync.WaitGroup) {
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

		query := bson.M{"time": bson.M{"$eq": 0}}

		if tx1stSynced {
			// skipt block 0
			query = bson.M{
				"blockNumber": bson.M{"$ne": 0},
				"time":        bson.M{"$eq": 0}}
		}

		err = txnC.Find(query).
			Sort("_id").Limit(step).All(&txns)

		if err != nil {
			log.Println("UpdateTxns query txns collection error:", err)
			continue
		}

		var wg sync.WaitGroup

		//length := len(txns)
		//p := length / 10
		//wg.Add(10)
		//
		//for index := 0; index < 10; index++ {
		//	if index == 9 {
		//		go func(start int, end int) {
		//			defer wg.Done()
		//			for i := start; i < end; i++ {
		//				txn := txns[i]
		//				newTxn, err := db.RpcGetTxByHash(txn.Hash)
		//				if err != nil {
		//					log.Println("UpdateTxns RpcGetTxByHash error:", err)
		//					return
		//				}
		//
		//				newTxn.BlockNumber = txn.BlockNumber
		//
		//				flatxns := newTxn.ToFlatTx()
		//				var tmpFlatxn db.FlatTx
		//
		//				for _, tx := range flatxns {
		//					errFind := flatxnC.Find(bson.M{"actionIndex": tx.ActionIndex,
		//						"hash": tx.Hash}).One(&tmpFlatxn)
		//
		//					// flatxn not found
		//					if errFind != nil {
		//						errInsert := flatxnC.Insert(*tx)
		//
		//						if errInsert != nil {
		//							log.Println("failed to insert to flatxnC")
		//							//	TODO: save those record to database to try again
		//							continue
		//						}
		//						log.Println("Insert flatxn on block", tx.BlockNumber)
		//					}
		//				}
		//
		//				errUpdate := txnC.Update(bson.M{"hash": newTxn.Hash},
		//					bson.M{
		//						"$set": bson.M{
		//							"time":       newTxn.Time,
		//							"expiration": newTxn.Expiration,
		//							"gasPrice":   newTxn.GasPrice,
		//							"gasLimit":   newTxn.GasLimit,
		//							"actions":    newTxn.Actions,
		//							"signers":    newTxn.Signers,
		//							"signs":      newTxn.Signs,
		//							"publisher":  newTxn.Publisher}})
		//
		//				if errUpdate != nil {
		//					log.Println("failed to update txn")
		//					//	TODO: save failed record to database to try again
		//					return
		//				}
		//
		//				if newTxn.BlockNumber == 0 {
		//					tx1stSynced = true
		//				}
		//				log.Println("Update Txn on block", newTxn.BlockNumber)
		//			}
		//		}(index*9, length)
		//	} else {
		//		go func(start int, end int) {
		//			defer wg.Done()
		//			for i := start; i < end; i++ {
		//				txn := txns[i]
		//				newTxn, err := db.RpcGetTxByHash(txn.Hash)
		//				if err != nil {
		//					log.Println("UpdateTxns RpcGetTxByHash error:", err)
		//					return
		//				}
		//
		//				newTxn.BlockNumber = txn.BlockNumber
		//
		//				flatxns := newTxn.ToFlatTx()
		//				var tmpFlatxn db.FlatTx
		//
		//				for _, tx := range flatxns {
		//					errFind := flatxnC.Find(bson.M{"actionIndex": tx.ActionIndex,
		//						"hash": tx.Hash}).One(&tmpFlatxn)
		//
		//					// flatxn not found
		//					if errFind != nil {
		//						errInsert := flatxnC.Insert(*tx)
		//
		//						if errInsert != nil {
		//							log.Println("failed to insert to flatxnC")
		//							//	TODO: save those record to database to try again
		//							continue
		//						}
		//						log.Println("Insert flatxn on block", tx.BlockNumber)
		//					}
		//				}
		//
		//				errUpdate := txnC.Update(bson.M{"hash": newTxn.Hash},
		//					bson.M{
		//						"$set": bson.M{
		//							"time":       newTxn.Time,
		//							"expiration": newTxn.Expiration,
		//							"gasPrice":   newTxn.GasPrice,
		//							"gasLimit":   newTxn.GasLimit,
		//							"actions":    newTxn.Actions,
		//							"signers":    newTxn.Signers,
		//							"signs":      newTxn.Signs,
		//							"publisher":  newTxn.Publisher}})
		//
		//				if errUpdate != nil {
		//					log.Println("failed to update txn")
		//					//	TODO: save failed record to database to try again
		//					return
		//				}
		//
		//				if newTxn.BlockNumber == 0 {
		//					tx1stSynced = true
		//				}
		//				log.Println("Update Txn on block", newTxn.BlockNumber)
		//			}
		//		}(index, index*p)
		//	}
		//}
		//wg.Wait()

		//for index := 0; index < length; index += 10 {
		//	if length-index < 10 {
		//		wg.Add(length - index)
		//	} else {
		//		wg.Add(10)
		//	}
		//
		//	for cur := index; cur < index+10 && cur < length; cur++ {
		//		txn := txns[cur]
		//		go updateTxns(&tx1stSynced, txn, flatxnC, txnC, &wg)
		//	}
		//	wg.Wait()
		//}

		for _, txn := range txns {
			wg.Add(1)
			go updateTxns(&tx1stSynced, txn, flatxnC, txnC, &wg)
		}
		wg.Wait()
		log.Println("Update txns goroutine finish, txn length: ", len(txns))
	}
}

func updateTxns(tx1stSynced *bool, txn *db.Tx, flatxnC *mgo.Collection, txnC *mgo.Collection, wg *sync.WaitGroup) {
	defer wg.Done()
	newTxn, err := db.RpcGetTxByHash(txn.Hash)

	if err != nil {
		log.Println("UpdateTxns RpcGetTxByHash error:", err)
		return
	}

	newTxn.BlockNumber = txn.BlockNumber

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
				continue
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
		return
	}

	if newTxn.BlockNumber == 0 {
		*tx1stSynced = true
	}
}
