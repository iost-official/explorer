package db

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/spf13/viper"
	"log"
)

var (
	MongoLink = "mongodb://127.0.0.1:27017"
	Db        string
)

const (
	CollectionBlocks     = "blocks"
	CollectionTmpTxs     = "tmpTxs"
	CollectionTxs        = "txs"
	CollectionRpcTxs     = "rpcErrTxs"
	CollectionFlatTx     = "flatxs"
	CollectionFBlocks    = "fBlocks"
	CollectionAccount    = "accounts"
	CollectionTaskCursor = "taskCursors"
	CollectionBlockPay   = "blockPays"
	CollectionApplyIost  = "applyTestIOST"
)

func InitConfig() {
	dbConfig := viper.GetStringMapString("mongodb")
	Db = dbConfig["db"]
	MongoLink = fmt.Sprintf("mongodb://%s:%s", dbConfig["host"], dbConfig["port"])
	fmt.Println("mongolink", Db, MongoLink)

	// create index
	col, err := GetCollection(CollectionFlatTx)
	if err != nil {
		log.Fatalln("Flat collection create index, get collection error", err)
	}
	err = col.EnsureIndexKey("from")
	err = col.EnsureIndexKey("to")
	err = col.EnsureIndexKey("publisher")
	err = col.EnsureIndexKey("hash")
	err = col.EnsureIndexKey("blockNumber")
	if err != nil {
		log.Fatalln("Flat collection create index error", err)
	}

	colTx, err := GetCollection(CollectionTxs)
	if err != nil {
		log.Fatalln("Flat collection create index, get collection error", err)
	}
	err = colTx.EnsureIndexKey("hash")
	err = colTx.EnsureIndexKey("blockNumber")
	err = colTx.EnsureIndexKey("mark")
	err = colTx.EnsureIndexKey("time")
	err = colTx.EnsureIndexKey("externalId")
	colBlock, err := GetCollection(CollectionBlocks)
	if err != nil {
		log.Fatalln("Block collection create index, get collection error", err)
	}
	err = colBlock.EnsureIndexKey("hash")
	err = colBlock.EnsureIndexKey("blockNumber")

	colTmpTx, err := GetCollection(CollectionTmpTxs)
	if err != nil {
		log.Fatalln("get tmp tx collection failed")
	}
	colTmpTx.EnsureIndexKey("blockNumber")
}

func EnsureCapped() {
	gb := int64(1 << 30)
	ensureCapped(CollectionFlatTx, 4*gb)
	ensureCapped(CollectionTxs, 4*gb)
	ensureCapped(CollectionTmpTxs, 2*gb)
}

func ensureCapped(col string, size int64) {
	var doc bson.M
	db, err := GetDb()
	if nil != err {
		log.Fatalln("ensure capped get db error", err)
		return
	}
	err = db.Run(bson.D{{Name: "collStats", Value: col}}, &doc)
	if nil != err {
		log.Fatalln("ensure capped get coll Stats error", err)
	}
	capped, ok := doc["capped"].(bool)
	log.Println("is ok", ok)
	if ok && !capped {
		err := db.Run(bson.D{
			{Name: "convertToCapped", Value: col},
			{Name: "size", Value: size}}, &doc)
		if nil != err {
			log.Println("convert to capped error", err)
		} else {
			log.Println("convert to capped result", doc)
		}
	}
}
