package db

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var (
	MongoLink = "mongodb://127.0.0.1:27017"
	Db string
)

const (
	CollectionBlocks  = "blocks"
	CollectionTxs     = "txs"
	CollectionFlatTx  = "flatxs"
	CollectionFBlocks = "fBlocks"
	CollectionAccount = "accounts"
	CollectionTaskCursor = "taskCursors"
	CollectionBlockPay = "blockPays"
	CollectionApplyIost = "applyTestIOST"
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
	colBlock, err := GetCollection(CollectionBlocks)
	if err != nil {
		log.Fatalln("Block collection create index, get collection error", err)
	}
	err = colBlock.EnsureIndexKey("hash")
	err = colBlock.EnsureIndexKey("blockNumber")
}


