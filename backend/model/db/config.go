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

	// create index
	col, err := GetCollection(CollectionFlatTx)
	if err != nil {
		log.Fatalln("Flat collection create index, get collection error", err)
	}
	err = col.EnsureIndexKey("from")
	err = col.EnsureIndexKey("to")
	err = col.EnsureIndexKey("publisher")
	err = col.EnsureIndexKey("hash")
	if err != nil {
		log.Fatalln("Flat collection create index error", err)
	}
}


