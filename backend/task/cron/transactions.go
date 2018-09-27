package cron

import (
	"github.com/globalsign/mgo/bson"
	"github.com/iost-official/explorer/backend/model/db"
	"log"
	"sync"
	"time"
)

func UpdateTxns(wg *sync.WaitGroup, mark int) {
	defer wg.Done()

	ticker := time.NewTicker(time.Second)

	txnC, err := db.GetCollection(db.CollectionTxs)

	if err != nil {
		log.Println("UpdateTxns get txns collection error:", err)
		return
	}

	tmpTxc, err := db.GetCollection(db.CollectionTmpTxs)

	if nil != err {
		log.Println("Update txns get tmp txs collection error:", err)
		return
	}

	flatxnC, err := db.GetCollection(db.CollectionFlatTx)

	if err != nil {
		log.Println("UpdateTxns get flatxs collection error:", err)
		return
	}

	rpcErrtxc, err := db.GetCollection(db.CollectionRpcTxs)

	if nil != err {
		log.Println("Update txns get rpc err collection error:", err)
		return
	}

	var startExternalId bson.ObjectId = ""

	for range ticker.C {
		step := 2000
		var txns = make([]*db.TmpTx, 0)

		var txn db.Tx
		var query = bson.M{"mark": mark}
		if "" == startExternalId {
			err = txnC.Find(bson.M{"mark": mark}).Sort("-externalId").Limit(1).One(&txn)
			if nil != err {
				if err.Error() != "not found" {
					log.Println("update tmpTx query error:", err)
					continue
				}
			} else {
				startExternalId = txn.ExternalId
				query = bson.M{"_id": bson.M{"$gt": startExternalId},
					"mark": mark}
			}
		} else {
			query = bson.M{"_id": bson.M{"$gt": startExternalId},
				"mark": mark}
		}

		err = tmpTxc.Find(query).Sort("_id").Limit(step).All(&txns)
		if nil != err {
			log.Println("query tmp txs error", err)
			continue
		}

		log.Println("need update txs length:", len(txns))

		var flatxs []interface{}
		var txs []interface{}
		var rpcErrTxs []interface{}

		for _, tmpTx := range txns {

			newTxn, err := db.RpcGetTxByHash(tmpTx.Hash)

			if err != nil {
				log.Println("UpdateTxns RpcGetTxByHash error:", err)
				tx := db.Tx{BlockNumber: tmpTx.BlockNumber, Hash: tmpTx.Hash, Mark: tmpTx.Mark, ExternalId: tmpTx.Id}
				rpcErrTxs = append(rpcErrTxs, tx)
				continue
			}
			newTxn.BlockNumber = tmpTx.BlockNumber

			flatxns := newTxn.ToFlatTx()

			for _, tx := range flatxns {
				flatxs = append(flatxs, *tx)
			}

			newTxn.Mark = tmpTx.Mark
			newTxn.ExternalId = tmpTx.Id
			startExternalId = newTxn.ExternalId
			txs = append(txs, *newTxn)
		}

		if len(txs) != 0 {
			err := txnC.Insert(txs...)
			if nil != err {
				log.Println("fail to insert txs, err: ", err)
			} else {
				log.Println("update txs, size: ", len(txs))
			}
		}

		if len(rpcErrTxs) != 0 {
			err := rpcErrtxc.Insert(rpcErrTxs...)
			if nil != err {
				log.Println("fail to insert rpc err txs, err: ", err)
			}
		}

		if len(flatxs) != 0 {
			err := flatxnC.Insert(flatxs...)
			if nil != err {
				log.Println("fail to insert flatxs, err: ", err)
			} else {
				log.Println("update flatxs, size: ", len(flatxs))
			}
		}
	}
}

func UpdateRpcErrTxns(wg *sync.WaitGroup) {
	wg.Done()
	ticker := time.NewTicker(time.Second * 2)

	rpcErrtxc, err := db.GetCollection(db.CollectionRpcTxs)

	if nil != err {
		log.Println("Update rpc err txns get rpc err collection error:", err)
		return
	}

	flatxnC, err := db.GetCollection(db.CollectionFlatTx)

	if err != nil {
		log.Println("UpdateTxns get flatxs collection error:", err)
		return
	}

	txnC, err := db.GetCollection(db.CollectionTxs)

	if err != nil {
		log.Println("UpdateTxns get txns collection error:", err)
		return
	}

	for range ticker.C {
		step := 2000
		var txns = make([]*db.Tx, 0)

		err := rpcErrtxc.Find(nil).Limit(step).All(&txns)

		if nil != err {
			if err.Error() != "not found" {
				log.Println("Update rpc err txs error", err)
			}
			continue
		}

		var flatxs []interface{}
		var txs []interface{}

		for _, txn := range txns {

			var exits db.Tx

			err := txnC.Find(bson.M{"_id": txn.ExternalId}).Limit(1).One(&exits)

			if nil == err {
				log.Println("txs already synced")
				err := rpcErrtxc.Remove(bson.M{"externalId": txn.ExternalId})
				if nil != err {
					log.Println("fail to remove rpc err txs, error:", err)
				}
				continue
			}

			newTxn, err := db.RpcGetTxByHash(txn.Hash)

			if err != nil {
				log.Println("UpdateTxns RpcGetTxByHash error:", err)
				continue
			}

			flatxns := newTxn.ToFlatTx()

			for _, tx := range flatxns {
				flatxs = append(flatxs, *tx)
			}

			newTxn.BlockNumber = txn.BlockNumber
			newTxn.ExternalId = txn.ExternalId
			newTxn.Mark = txn.Mark
			txs = append(txs, *newTxn)

			er := rpcErrtxc.Remove(bson.M{"externalId": txn.ExternalId})

			if nil != er {
				log.Println("remove rpc err tx error:", er)
			}
		}

		if len(txs) != 0 {
			err := txnC.Insert(txs...)
			if nil != err {
				log.Println("fail to insert txs, err: ", err)
			} else {
				log.Println("update rpc err txs, size: ", len(txs))
			}
		}

		if len(flatxs) != 0 {
			err := flatxnC.Insert(flatxs...)
			if nil != err {
				log.Println("fail to insert flatxs, err: ", err)
			} else {
				log.Println("update rpc err flatxs, size: ", len(flatxs))
			}
		}
	}

}
