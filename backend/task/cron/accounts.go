package cron

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/iost-official/explorer/backend/model/blkchain"
	"github.com/iost-official/explorer/backend/model/db"
	"log"
	"sync"
	"time"
)

func UpdateAccounts(wg *sync.WaitGroup) {
	defer wg.Done()

	flatTxCol, err := db.GetCollection(db.CollectionFlatTx)
	if err != nil {
		log.Fatalln("Flat Collection get failed")
	}

	accountCol, err := db.GetCollection(db.CollectionAccount)
	if err != nil {
		log.Fatalln("Account collection get failed")
	}

	ticker := time.NewTicker(time.Second * 10)
	for _ = range ticker.C {
		fmt.Println("Update account info")
		var query = bson.M{"actionName": "Transfer"}
		cursor, err := db.GetAccountTaskCursor()
		if err == nil {
			query["_id"] = bson.M{"$gt": cursor}
		}
		var txs []db.FlatTx
		err = flatTxCol.Find(query).Sort("_id").Limit(50).All(&txs)
		for _, ft := range txs {
			// ===== update from account
			fromB, err := blkchain.GetBalance(ft.From)
			if err != nil {
				fmt.Println("Get balance failed", err)
			}
			fromCount, err := db.GetAccountTxCount(ft.From)
			if err != nil {
				fmt.Println("Get txcount error", err)
			}
			_, err = accountCol.Upsert(bson.M{"address": ft.From}, bson.M{"$set": bson.M{"balance": fromB, "tx_count": fromCount}}) // TODO: check error update
			if err != nil {
				fmt.Println("Update failed", err)
			}
			// ====== update to account
			toB, err := blkchain.GetBalance(ft.To)
			if err != nil {
				fmt.Println("Get balance failed", err)
			}
			toCount, err := db.GetAccountTxCount(ft.To)
			if err != nil {
				fmt.Println("Get tx failed", err)
			}
			_, err = accountCol.Upsert(bson.M{"address": ft.To}, bson.M{"$set": bson.M{"balance": toB, "tx_count": toCount}})
			if err != nil {
				fmt.Println("Update failed", err)
			}
		}

		if len(txs) > 0 {
			err = db.UpdateAccountTaskCursor(txs[len(txs)-1].Id)
			if err != nil {
				fmt.Println("Update cursor error: ", err)
			}
		}
	}
}
