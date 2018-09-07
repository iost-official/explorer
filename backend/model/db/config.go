package db

import (
	"fmt"
	"github.com/spf13/viper"
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
}


