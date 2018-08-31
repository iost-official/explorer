package cron

import (
	"github.com/iost-official/explorer/backend/model/blkchain"
	"github.com/iost-official/explorer/backend/model/dbv2"
	"log"
	"sync"
	"time"
)

func UpdateBlocks2(ws *sync.WaitGroup) {
	defer ws.Done()

	collection, err := dbv2.GetCollection("blocks")
	if nil != err {
		log.Println("can not get blocks collection when update", err)
		return
	}

	txCollection, err := dbv2.GetCollection("txs")

	if nil != err {
		log.Println("can not get txs collection when update", err)
	}

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
				continue
			}

			err = collection.Insert(block)

			if err != nil {
				log.Println("updateBlock2 insert mongo error:", err)
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
				}
			}

			insertLen++
			log.Println("updateBlock insert mongo height:", topHeightInMongo)
		}

		log.Println("updateBlock inserted len:", insertLen)

	}
}
