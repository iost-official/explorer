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
		var query = bson.M{
			//"actionName": "Transfer",
		}
		cursor, err := db.GetAccountTaskCursor()
		if err == nil {
			query["_id"] = bson.M{"$gt": cursor}
		}
		var txs []db.FlatTx
		err = flatTxCol.Find(query).Sort("_id").Limit(50).All(&txs)
		for _, ft := range txs {
			// ===== update from account
			var fromB int64
			if ft.From[0:4] == "IOST" {  // IOST 地址才会获取
				fromB, err = blkchain.GetBalance(ft.From)
				if err != nil {
					fmt.Println("Get balance failed", err)
				}
			}
			_, err = accountCol.Upsert(bson.M{"address": ft.From}, bson.M{"$set": bson.M{"balance": fromB}})
			if err != nil {
				fmt.Println("Update failed", err)
			}

			if ft.To[0:5] == "iost." {  // 跳过特殊地址
				continue
			}

			var toB int64
			// ====== update to account
			if ft.To[0:4] == "IOST" {
				toB, err = blkchain.GetBalance(ft.To)
				if err != nil {
					fmt.Println("Get balance failed", err)
				}
			}
			_, err = accountCol.Upsert(bson.M{"address": ft.To}, bson.M{"$set": bson.M{"balance": toB}})
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
