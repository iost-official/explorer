package cron

import (
	"github.com/iost-official/explorer/backend/model/blkchain"
	"github.com/iost-official/explorer/backend/model/dbv2"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"sync"
	"time"
)

func UpdateBlocks2(ws *sync.WaitGroup) {
	defer ws.Done()

	collection, err := dbv2.GetCollection(dbv2.CollectionBlocks)
	if nil != err {
		log.Println("can not get blocks collection when update", err)
		return
	}

	txCollection, err := dbv2.GetCollection(dbv2.CollectionTxs)

	if nil != err {
		log.Println("can not get txs collection when update", err)
		return
	}

	//record sync failed blocks
	fBlockCollection, err := dbv2.GetCollection(dbv2.CollectionFBlocks)

	ticker := time.NewTicker(time.Second * 5)

	for range ticker.C {
		var topHeightInChain int64 = 0
		var topHeightInMongo int64 = 0

		topBlcHeight, err := blkchain.GetCurrentBlockHeight()

		if err != nil {
			log.Println("updateBlock get topBlk in chain error:", err)
			continue
		}

		topHeightInChain = topBlcHeight

		topBlkInMongo, err := dbv2.GetTopBlock2()

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

			block, txHashes, err := dbv2.GetBlockInfoByNum(topHeightInMongo)

			if nil != err {
				log.Println("UpdateBlock2 GetBlockInfoByNum error", err)

				err := recordFailedUpdateBlock(topHeightInMongo, fBlockCollection)
				if nil != err {
					log.Println("UpdateBlock2 record sync failed block error", err)
				}
				continue
			}

			err = collection.Insert(block)

			if err != nil {
				log.Println("updateBlock2 insert mongo error:", err)

				err := recordFailedUpdateBlock(topHeightInMongo, fBlockCollection)
				if nil != err {
					log.Println("UpdateBlock2 record sync failed block error", err)
				}

				continue
			}

			if nil != txHashes {
				txs := make([]interface{}, len(*txHashes))
				for index, txHash := range *txHashes {
					txs[index] = dbv2.Tx{
						Hash:        txHash,
						BlockNumber: topHeightInMongo,
					}
				}
				err := txCollection.Insert(txs...)
				if nil != err {
					log.Println("UpdateBlock2 insert txs error", err)
					err := recordFailedUpdateBlock(topHeightInMongo, fBlockCollection)

					if nil != err {
						// fix it?
						log.Println("UpdateBlock2 Record failed insert error", err)
					}
				}
			}

			insertLen++
			log.Println("updateBlock insert mongo height:", topHeightInMongo)
		}

		log.Println("updateBlock inserted len:", insertLen)

	}
}

func ProcessFailedSyncBlocks(ws *sync.WaitGroup) {
	defer ws.Done()

	log.Println("start process f block task")

	collection, err := dbv2.GetCollection(dbv2.CollectionBlocks)
	if nil != err {
		log.Println("can not get blocks collection when update", err)
		return
	}

	fBlockCollection, err := dbv2.GetCollection(dbv2.CollectionFBlocks)
	if nil != err {
		log.Println("Process failed sync blocks get f blocks collection error", err)
		return
	}

	txCollection, err := dbv2.GetCollection(dbv2.CollectionTxs)

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

	ticker := time.NewTicker(time.Second * 5)

	for range ticker.C {
		var fBlockList = make([]*dbv2.FailBlock, 0)
		fBlockCollection.Find(query).Sort("blockNumber").All(&fBlockList)
		log.Println("remain to process block count", len(fBlockList))
		for _, fBlock := range fBlockList {
			block, txHashes, err := dbv2.GetBlockInfoByNum(fBlock.BlockNumber)

			if err != nil {
				log.Println("Process failed blocks rpc call error:", err)
				fBlock.RetryTimes++
				fBlockCollection.Update(bson.M{"blockNumber": fBlock.BlockNumber}, bson.M{"retryTimes": fBlock.RetryTimes})
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
					fBlockCollection.Update(bson.M{"blockNumber": fBlock.BlockNumber}, bson.M{"retryTimes": fBlock.RetryTimes})
					continue
				}
			}

			if nil != txHashes {
				txs := make([]interface{}, len(*txHashes))
				for index, txHash := range *txHashes {
					txs[index] = dbv2.Tx{
						Hash:        txHash,
						BlockNumber: fBlock.BlockNumber,
					}
				}
				err := txCollection.Insert(txs...)
				if nil != err {
					fBlock.RetryTimes++
					fBlockCollection.Update(bson.M{"blockNumber": fBlock.BlockNumber}, bson.M{"retryTimes": fBlock.RetryTimes})
					log.Println("UpdateBlock2 insert txs error", err)
					continue
				}
			}
			fBlockCollection.Update(bson.M{"blockNumber": fBlock.BlockNumber}, bson.M{"processed": true})
		}
	}
}

func recordFailedUpdateBlock(blockNumber int64, fBlockCollection *mgo.Collection) error {
	fBlock := dbv2.FailBlock{
		BlockNumber: blockNumber,
		RetryTimes:  0,
		Processed:   false,
	}

	err := fBlockCollection.Insert(fBlock)
	return err
}
