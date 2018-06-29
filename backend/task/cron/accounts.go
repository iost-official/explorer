package cron

import (
	"sync"
	"explorer/model/db"
	"log"
	"gopkg.in/mgo.v2/bson"
	"time"
	"explorer/model/blockchain"
)

func UpdateAccounts(wg *sync.WaitGroup) {
	defer wg.Done()

	txnDC, err := db.GetCollection("txnsdetail")
	if err != nil {
		log.Println("UpdateAccounts get collection error:", err)
		return
	}

	accountC, err := db.GetCollection("accounts")
	if err != nil {
		log.Println("UpdateAccounts get collection error:", err)
		return
	}

	queryFrom := bson.M{
		"from": bson.M{
			"$ne": "",
		},
	}
	queryTo := bson.M{
		"to": bson.M{
			"$ne": "",
		},
	}

	ticker := time.NewTicker(time.Second * 60)
	for _ = range ticker.C {
		fromList := make([]string, 0)
		err = txnDC.Find(queryFrom).Distinct("from", &fromList)
		if err != nil {
			log.Println("UpdateAccounts find from error:", err)
			continue
		}

		toList := make([]string, 0)
		err = txnDC.Find(queryTo).Distinct("to", &toList)
		if err != nil {
			log.Println("UpdateAccounts find to error:", err)
		}

		fromList = append(fromList, toList...)

		for _, account := range fromList {
			balance, err := blockchain.GetBalanceByKey(account)
			if err != nil {
				log.Println("UpdateAccounts GetBalanceByKey error:", err)
			}

			txnLen, err := db.GetTxnDetailLenByAccount(account)
			if err != nil {
				log.Println("UpdateAccounts GetTxnDetailLenByAccount error:", err)
			}

			selector := bson.M{
				"address": account,
			}
			accountC.Upsert(selector, &db.Account{
				Address: account,
				Balance: balance,
				TxCount: txnLen,
			})
			log.Println("UpdateAccounts address:", account, "amount:", balance, "updated.")
		}
	}
}
