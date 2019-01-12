package db

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	MongoLink     = "mongodb://127.0.0.1:27017"
	MongoUser     = ""
	MongoPassWord = ""
	Db            string
)

const (
	CollectionBlocks        = "blocks"
	CollectionTxs           = "txs"
	CollectionFlatTx        = "flatxs"
	CollectionAccount       = "accounts"
	CollectionAccountTx     = "accountTx"
	CollectionAccountPubkey = "accountPubkey"
	CollectionContract      = "contracts"
	CollectionContractTx    = "contractTx"
	CollectionTaskCursor    = "taskCursors"
	CollectionBlockPay      = "blockPays"
	CollectionApplyIOST     = "applyIOST"
)

func InitConfig() {
	dbConfig := viper.GetStringMapString("mongodb")
	Db = dbConfig["db"]
	MongoUser = dbConfig["username"]
	MongoPassWord = dbConfig["password"]
	//MongoLink = fmt.Sprintf("%s:%s", dbConfig["host"], dbConfig["port"])
	MongoLink = dbConfig["mongoLink"]
	fmt.Println("mongolink", Db, MongoLink)
}
