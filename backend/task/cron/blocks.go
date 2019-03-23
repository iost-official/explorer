package cron

import (
	"log"
	"sync"
	"time"

	"github.com/iost-official/explorer/backend/model/blockchain"
	"github.com/iost-official/explorer/backend/model/blockchain/rpcpb"
	"github.com/iost-official/explorer/backend/model/db"
)

func UpdateBlocks(ws *sync.WaitGroup) {
	defer ws.Done()

	blockChannel := make(chan *rpcpb.Block, 10)
	go insertBlock(blockChannel)

	ticker := time.NewTicker(time.Second)

	var topHeightInMongo int64
	for range ticker.C {
		topBlkInMongo, err := db.GetTopBlock()
		if err != nil {
			log.Println("updateBlock get topBlk in mongo error:", err)
			if err.Error() != "not found" {
				continue
			} else {
				topHeightInMongo = 0
				break
			}
		}
		topHeightInMongo = topBlkInMongo.Number + 1
		log.Println("Got Top Block From Mongo With Number: ", topHeightInMongo)
		break
	}

	for {
		wg := new(sync.WaitGroup)
		wg.Add(10)
		for i := 0; i < 10; i++ {
			go func(topHeightInMongo int64) {
				for {
					blockRspn, err := blockchain.GetBlockByNum(topHeightInMongo, true)
					if err != nil {
						log.Println("Download block", topHeightInMongo, "error:", err)
						time.Sleep(time.Second)
						continue
					}
					if blockRspn.Status == rpcpb.BlockResponse_PENDING {
						log.Println("Download block", topHeightInMongo, "Pending")
						time.Sleep(time.Second)
						continue
					}
					blockChannel <- blockRspn.Block
					log.Println("Download block", topHeightInMongo, " Succ!")
					wg.Done()
					return
				}
			}(topHeightInMongo + int64(i))
		}
		wg.Wait()
		topHeightInMongo += 10
	}
}

func insertBlock(blockChannel chan *rpcpb.Block) {
	collection := db.GetCollection(db.CollectionBlocks)

	for {
		select {
		case b := <-blockChannel:
			txs := b.Transactions

			wg := new(sync.WaitGroup)
			wg.Add(2)
			go func() {
				db.ProcessTxs(txs, b.Number)
				wg.Done()
			}()
			go func() {
				db.ProcessTxsForAccount(txs, b.Time, b.Number)
				wg.Done()
			}()
			wg.Wait()

			b.Transactions = make([]*rpcpb.Transaction, 0)
			err := collection.Insert(*b)

			if err != nil {
				log.Println("updateBlock insert mongo error:", err)
			}
		default:

		}
	}
}
