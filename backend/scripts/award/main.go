package main

import (
	"fmt"
	"github.com/iost-official/explorer/backend/model/db"
	"github.com/iost-official/explorer/backend/util/common"
)

func main() {
	db.Db = "mainnet"
	db.MongoLink = ""
	aid := "15542208000000000001555516800000000000100000"
	ainfo, err := db.GetProducerLevelInfo(aid)
	if err != nil {
		panic(err)
	}
	fmt.Println("%+v", ainfo)

	err = common.CalculateProducerContributions(ainfo)
	if err != nil {
		panic(err)
	}
}
