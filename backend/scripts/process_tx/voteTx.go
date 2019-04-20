package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/iost-official/explorer/backend/config"
	"github.com/iost-official/explorer/backend/model/blockchain/rpcpb"
	"github.com/iost-official/explorer/backend/model/db"
)

func retryWriteMongo(b *mgo.Bulk) {
	var retryTime int
	for {
		if _, err := b.Run(); err != nil {
			log.Println("fail to write data to mongo ", err)
			time.Sleep(time.Second)
			retryTime++
			if retryTime > 10 {
				log.Fatalln("fail to write data to mongo, retry time exceeds")
			}
			continue
		}
		return
	}
}

func ProcessTxs() {

	c := db.GetCollection(db.CollectionTxs)

	iter := c.Find(bson.M{}).Iter()
	var txStore db.TxStore
	for iter.Next(&txStore) {
		ProcessTxsForVote(txStore)
	}
	if err := iter.Close(); err != nil {
		fmt.Println(err.Error)
		return
	}
}

func main() {
	config.ReadConfig("")
	ProcessTxs()
}

func ProcessTxsForVote(txStore db.TxStore) {
	blockNumber := txStore.BlockNumber
	fmt.Println("ProcessingBlock: ", blockNumber)
	t := txStore.Tx
	voteTxC := db.GetCollection(db.CollectionVoteTx)
	voteTxB := voteTxC.Bulk()

	isVote := false

	for _, r := range t.TxReceipt.Receipts {
		if strings.Contains(r.FuncName, "vote_producer.iost") {
			isVote = true
		}
	}
	for _, a := range t.Actions {
		if a.Contract == "vote_producer.iost" {
			isVote = true
		}
	}

	if isVote {
		if t.TxReceipt.StatusCode == rpcpb.TxReceipt_SUCCESS {
			voteTxB.Upsert(bson.M{"txReceipt.txhash": t.TxReceipt.TxHash}, &db.VoteTx{t.Actions, t.TxReceipt, blockNumber})
		}
	}

	retryWriteMongo(voteTxB)
}
