package cron

import (
	"fmt"
	"github.com/iost-official/explorer/backend/model/blkchain"
	"github.com/iost-official/explorer/backend/model/db"
	"log"
	"sync"
	"time"
)

func CheckRestart(ws *sync.WaitGroup) {
	defer ws.Done()

	ticker := time.NewTicker(time.Second * 2)

	for range ticker.C {
		topHeightInChain, err := blkchain.GetCurrentBlockHeight()
		if err != nil {
			log.Println("updateBlock get topBlk in chain error:", err)
			continue
		}


		topBlkInMongo, err := db.GetTopBlock()

		if err != nil {
			log.Println("updateBlock get topBlk in mongo error:", err)
			continue
		}

		topHeightInMongo := topBlkInMongo.BlockNumber
		if topHeightInChain < topHeightInMongo {
			// TODO: CHECK
			fmt.Println("drop database")
		}

	}
}