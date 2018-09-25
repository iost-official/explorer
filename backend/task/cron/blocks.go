package cron

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/iost-official/explorer/backend/model/blkchain"
	"github.com/iost-official/explorer/backend/model/db"
	"log"
	"sync"
	"time"
)

func UpdateBlocks(ws *sync.WaitGroup) {
	defer ws.Done()

	collection, err := db.GetCollection(db.CollectionBlocks)
	if nil != err {
		log.Println("can not get blocks collection when update", err)
		return
	}

	tmpTxCollection, err := db.GetCollection(db.CollectionTmpTxs)

	if nil != err {
		log.Println("can not get txs collection when update", err)
		return
	}

	//record sync failed blocks
	fBlockCollection, err := db.GetCollection(db.CollectionFBlocks)

	ticker := time.NewTicker(time.Second * 2)

	for range ticker.C {
		var topHeightInChain int64 = 0
		var topHeightInMongo int64 = 0

		topBlcHeight, err := blkchain.GetCurrentBlockHeight()

		if err != nil {
			log.Println("updateBlock get topBlk in chain error:", err)
			continue
		}

		topHeightInChain = topBlcHeight

		topBlkInMongo, err := db.GetTopBlock()

		if err != nil {
			log.Println("updateBlock get topBlk in mongo error:", err)
			if err.Error() != "not found" {
				continue
			}
		} else {
			topHeightInMongo = topBlkInMongo.BlockNumber + 1
		}
		var insertLen int
		for ; topHeightInMongo <= topHeightInChain; topHeightInMongo++ {
			block, txHashes, err := db.GetBlockInfoByNum(topHeightInMongo)

			if nil != err {
				log.Println("UpdateBlock GetBlockInfoByNum error", err)

				err := recordFailedUpdateBlock(topHeightInMongo, fBlockCollection)
				if nil != err {
					log.Println("UpdateBlock record sync failed block error", err)
				}
				continue
			}

			err = collection.Insert(block)

			if err != nil {
				log.Println("updateBlock insert mongo error:", err)

				err := recordFailedUpdateBlock(topHeightInMongo, fBlockCollection)
				if nil != err {
					log.Println("UpdateBlock record sync failed block error", err)
				}

				continue
			}

			if nil != txHashes {
				txs := make([]interface{}, len(*txHashes))
				for index, txHash := range *txHashes {
					txs[index] = db.TmpTx{
						Hash:        txHash,
						BlockNumber: topHeightInMongo,
						Mark:        topHeightInMongo % 2,
					}
				}
				err := tmpTxCollection.Insert(txs...)
				if nil != err {
					log.Println("UpdateBlock insert txs error", err)
					err := recordFailedUpdateBlock(topHeightInMongo, fBlockCollection)

					if nil != err {
						// fix it?
						log.Println("UpdateBlock Record failed insert error", err)
					}
				}
			}

			insertLen++
			log.Println("updateBlock insert mongo height:", topHeightInMongo)
		}

		log.Println("updateBlock inserted len: ======", insertLen)

	}
}

func ProcessFailedSyncBlocks(ws *sync.WaitGroup) {
	defer ws.Done()

	collection, err := db.GetCollection(db.CollectionBlocks)
	if nil != err {
		log.Println("can not get blocks collection when update", err)
		return
	}

	fBlockCollection, err := db.GetCollection(db.CollectionFBlocks)
	if nil != err {
		log.Println("Process failed sync blocks get f blocks collection error", err)
		return
	}

	tmpTxCollection, err := db.GetCollection(db.CollectionTmpTxs)

	if nil != err {
		log.Println("Process failed sync blocks get txs collection error", err)
		return
	}

	query := bson.M{
		"processed": false,
		"retryTimes": bson.M{
			"$lte": 5,
		},
	}

	ticker := time.NewTicker(time.Second * 2)

	for range ticker.C {
		var fBlockList = make([]*db.FailBlock, 0)
		fBlockCollection.Find(query).Sort("blockNumber").All(&fBlockList)
		for _, fBlock := range fBlockList {
			block, txHashes, err := db.GetBlockInfoByNum(fBlock.BlockNumber)

			if err != nil {
				log.Println("Process failed blocks rpc call error:", err)
				fBlock.RetryTimes++
				fBlockCollection.Update(bson.M{"blockNumber": fBlock.BlockNumber}, bson.M{"$set": bson.M{
					"retryTimes": fBlock.RetryTimes,
				}})
				continue
			}

			count, err := collection.Find(bson.M{"blockNumber": fBlock.BlockNumber}).Count()

			if nil != err {
				log.Println("Process failed blocks count error", err)
				continue
			}

			if count == 0 {
				err = collection.Insert(block)

				if nil != err {
					log.Println("Process failed blocks insert block error", err)
					fBlock.RetryTimes++
					fBlockCollection.Update(bson.M{"blockNumber": fBlock.BlockNumber}, bson.M{"$set": bson.M{
						"retryTimes": fBlock.RetryTimes,
					}})
					continue
				}
			}

			if nil != txHashes {
				txs := make([]interface{}, len(*txHashes))
				for index, txHash := range *txHashes {
					txs[index] = db.TmpTx{
						Hash:        txHash,
						BlockNumber: fBlock.BlockNumber,
						Mark:        fBlock.BlockNumber % 2,
					}
				}
				err := tmpTxCollection.Insert(txs...)
				if nil != err {
					fBlock.RetryTimes++
					fBlockCollection.Update(bson.M{"blockNumber": fBlock.BlockNumber}, bson.M{"$set": bson.M{
						"retryTimes": fBlock.RetryTimes,
					}})
					log.Println("UpdateBlock2 insert txs error", err)
					continue
				}
			}
			fBlockCollection.Update(bson.M{"blockNumber": fBlock.BlockNumber}, bson.M{"$set": bson.M{
				"processed": true,
			}})
		}
	}
}

func UpdateBlockPay(wg *sync.WaitGroup) {
	defer wg.Done()

	txnC, err := db.GetCollection(db.CollectionTxs)
	if err != nil {
		log.Println("UpdateBlockCost get collection error:", err)
		return
	}

	blkPC, err := db.GetCollection(db.CollectionBlockPay)
	if err != nil {
		log.Println("UpdateBlockCost get collection error:", err)
		return
	}

	ticker := time.NewTicker(time.Second * 2)
	for range ticker.C {
		var topHeightInPay int64 = 0
		topPay, err := db.GetTopBlockPay()
		if err != nil {
			if err.Error() != "not found" {
				continue
			}
		} else {
			topHeightInPay = topPay.Height + 1
		}

		queryPip := []bson.M{
			{
				"$match": bson.M{
					"blockNumber": bson.M{
						"$gte": topHeightInPay,
					},
					"time": bson.M{
						"$ne": 0,
					},
				},
			},
			{
				"$group": bson.M{
					"_id": "$blockNumber",
					"avggasprice": bson.M{
						"$avg": "$gasPrice",
					},
					"totalgaslimit": bson.M{
						"$sum": "$gasLimit",
					},
				},
			},
		}

		var payList []*db.BlockPay
		err = txnC.Pipe(queryPip).All(&payList)
		if err != nil {
			log.Println("UpdateBlockPay pipeline error:", err)
			continue
		}

		for _, pay := range payList {
			selector := bson.M{
				"_id": pay.Height,
			}
			blkPC.Upsert(selector, pay)
			log.Println("UpdateBlockPay block:", pay.Height, "inserted")
		}
	}
}

func recordFailedUpdateBlock(blockNumber int64, fBlockCollection *mgo.Collection) error {
	fBlock := db.FailBlock{
		BlockNumber: blockNumber,
		RetryTimes:  0,
		Processed:   false,
	}

	err := fBlockCollection.Insert(fBlock)
	return err
}
