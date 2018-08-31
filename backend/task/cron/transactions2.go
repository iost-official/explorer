package cron

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/iost-official/explorer/backend/model/db"
	"github.com/iost-official/explorer/backend/model/dbv2"
	"log"
	"sync"
	"time"
)

func UpdateTxns(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	ticker := time.NewTicker(time.Second * 5)
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
			fmt.Println(newTxn)

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

			fmt.Println(newTxn)
		}

		start += step
	}
}

//func UpdateTxs(wg *sync.WaitGroup) {
//	defer wg.Done()
//
//	blkC, err := dbv2.GetCollection("blocks")
//	if err != nil {
//		log.Println("UpdateTxns get block collection error:", err)
//		return
//	}
//
//	txnC, err := db.GetCollection("txns")
//	if err != nil {
//		log.Println("UpdateTxns get txns collection error:", err)
//		return
//	}
//
//	ticker := time.NewTicker(time.Second * 2)
//
//	for range ticker.C {
//		var (
//			topTxBlkHeight      int64 = 0
//			topBlkHeightInMongo int64 = 0
//		)
//		topTxInMongo, err := db.GetTopTxn()
//		if err != nil {
//			if err.Error() != "not found" {
//				continue
//			}
//		} else {
//			topTxBlkHeight = topTxInMongo.BlockHeight
//		}
//
//		topBlkInMongo, err := db.GetTopBlock()
//		if err != nil {
//			log.Println("UpdateTxs get topBlk in mongo error:", err)
//			continue
//		} else {
//			topBlkHeightInMongo = topBlkInMongo.Head.Number
//		}
//
//		if topBlkHeightInMongo == topTxBlkHeight {
//			log.Println("UpdateTxs txn height equals blk height, continue...")
//			continue
//		}
//
//		log.Printf("UpdateTxs topTxBlkHeight: %d topBlkHeightInMongo: %d\n", topTxBlkHeight, topBlkHeightInMongo)
//
//		blkList := make([]*rpc.BlockInfo, 0)
//		blkQuery := bson.M{
//			"head.number": bson.M{
//				"$gt": topTxBlkHeight,
//			},
//		}
//
//		err = blkC.Find(blkQuery).Limit(10).All(&blkList)
//		if err != nil {
//			log.Println("UpdateTxs getBlk error:", err)
//			continue
//		}
//
//		for _, blk := range blkList {
//			txnList := make([]interface{}, 0)
//			queryStart := time.Now()
//			for _, txnKey := range blk.TxList {
//				txn, err := blockchain.GetTxnByKey(txnKey)
//				if err != nil {
//					log.Println("UpdateTxns get tx error:", err)
//					continue
//				}
//				txnList = append(txnList, &db.ExplorerTx{
//					BlockHeight: blk.Head.Number,
//					Tx:          txn,
//					TxHash:      txn.Hash(),
//				})
//			}
//			if len(txnList) == 0 {
//				txn := &db.ExplorerTx{
//					BlockHeight: blk.Head.Number,
//					Tx:          nil,
//				}
//				txnList = append(txnList, txn)
//			}
//			queryTime := time.Since(queryStart)
//
//			insertStart := time.Now()
//			bulk := txnC.Bulk()
//			bulk.Unordered()
//			bulk.Insert(txnList...)
//			_, err = bulk.Run()
//			if err != nil {
//				log.Println("updateTxns blk:", blk.Head.Number, "txns:", len(txnList), "insert mongo error:", err)
//				continue
//			}
//			log.Println("updateTxns blk:", blk.Head.Number, "txns:", len(txnList), "insert success, time:", time.Since(insertStart), "search time:", queryTime)
//		}
//	}
//}

//func UpdateTxnDetail(wg *sync.WaitGroup) {
//	defer wg.Done()
//
//	txnC, err := db.GetCollection("txns")
//	if err != nil {
//		log.Println("UpdateTxnDetail get txns collection error:", err)
//		return
//	}
//
//	txnDC, err := db.GetCollection("txnsdetail")
//	if err != nil {
//		log.Println("UpdateTxnDetail get txns-detail collection error:", err)
//		return
//	}
//
//	txnQuery := bson.M{
//		"tx": bson.M{
//			"$ne": nil,
//		},
//	}
//
//	ticker := time.NewTicker(time.Second)
//	for _ = range ticker.C {
//		topTxnDetail, err := db.GetTopTxnDetail()
//		if err != nil && err.Error() != "not found" {
//			continue
//		}
//		if err == nil {
//			txnQuery["_id"] = bson.M{
//				"$gt": topTxnDetail.MgoSourceId,
//			}
//		}
//
//		var txnList []*db.ExplorerTx
//		err = txnC.Find(txnQuery).Sort("_id").Limit(5000).All(&txnList)
//		if err != nil {
//			log.Println("UpdateTxnDetail get txnList error:", err)
//			continue
//		}
//
//		if len(txnList) == 0 {
//			time.Sleep(time.Second * 2)
//			log.Println("UpdateTxnDetail details height equals txn, continue...")
//			continue
//		}
//
//		var mgoTxList []interface{}
//		for _, txn := range txnList {
//			mgoTx := txn.GenerateMgoTx()
//			if mgoTx == nil {
//				continue
//			}
//			mgoTxList = append(mgoTxList, mgoTx)
//		}
//
//		err = txnDC.Insert(mgoTxList...)
//		if err != nil {
//			log.Println("UpdateTxnDetail insert txnDC error:", err, "len:", len(mgoTxList))
//			continue
//		}
//		log.Println("UpdateTxnDetail insert txnDC success, len:", len(mgoTxList))
//	}
//}
