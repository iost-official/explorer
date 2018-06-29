package cron

import (
	"log"
	"sync"
	"time"

	"explorer/model/blockchain"
	"explorer/model/db"

	"gopkg.in/mgo.v2/bson"
)

func UpdateBlocks(wg *sync.WaitGroup) {
	defer wg.Done()

	collection, err := db.GetCollection("blocks")
	if err != nil {
		log.Println("updateBlock get collection error:", err)
		return
	}

	ticker := time.NewTicker(time.Second * 2)
	for _ = range ticker.C {
		var topHeightInChain int64 = 0
		var topHeightInMongo int64 = 0

		topBlkInChain, err := blockchain.GetTopBlock()
		if err != nil {
			log.Println("updateBlock get topBlk in chain error:", err)
			continue
		}
		topHeightInChain = topBlkInChain.Head.Number

		topBlkInMongo, err := db.GetTopBlock()
		if err != nil {
			log.Println("updateBlock get topBlk in mongo error:", err)
			if err.Error() != "not found" {
				continue
			}
		} else {
			topHeightInMongo = topBlkInMongo.Head.Number + 1
		}

		var insertLen int
		for ; topHeightInMongo <= topHeightInChain; topHeightInMongo++ {
			bInfo, err := blockchain.GetBlockByHeight(topHeightInMongo)
			if err != nil {
				log.Println("updateBlock getBlockByHeight error:", err)
				continue
			}

			err = collection.Insert(bInfo)
			if err != nil {
				log.Println("updateBlock insert mongo error:", err)
				continue
			}
			insertLen++
			log.Println("updateBlock insert mongo height:", topHeightInMongo)
		}

		log.Println("updateBlock inserted len:", insertLen)
	}
}

func UpdateBlockPay(wg *sync.WaitGroup)  {
	defer wg.Done()

	txnDC, err := db.GetCollection("txnsdetail")
	if err != nil {
		log.Println("UpdateBlockCost get collection error:", err)
		return
	}

	blkPC, err := db.GetCollection("blockpay")
	if err != nil {
		log.Println("UpdateBlockCost get collection error:", err)
		return
	}

	ticker := time.NewTicker(time.Second * 2)
	for _ = range ticker.C {
		var topHeightInPay int64 = 0
		topPay, err := db.GetTopBlockPay()
		if err != nil {
			if err.Error() != "not found" {
				continue
			}
		} else {
			topHeightInPay = topPay.Height
		}

		queryPip := []bson.M{
			bson.M{
				"$match": bson.M{
					"blockheight": bson.M{
						"$gte": topHeightInPay,
					},
				},
			},
			bson.M{
				"$group": bson.M{
					"_id": "$blockheight",
					"avggasprice": bson.M{
						"$avg": "$price",
					},
					"totalgaslimit": bson.M{
						"$sum": "$gaslimit",
					},
				},
			},
		}

		var payList []*db.BlockPay
		err = txnDC.Pipe(queryPip).All(&payList)
		if err != nil {
			log.Println("UpdateBlockPay pipline error:", err)
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
