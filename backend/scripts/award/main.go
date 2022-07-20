package main

import (
	"fmt"
	"github.com/GincoInc/iost-explorer/backend/model/db"
	"github.com/GincoInc/iost-explorer/backend/util/common"
)

func main() {
	db.Db = "mainnet"
	db.MongoLink = ""
	aid := "15542208000000000001555516800000000000100000"
	ainfo, err := db.GetProducerLevelInfo(aid)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", ainfo)

	err = common.CalculateProducerContributions(ainfo)
	if err != nil {
		panic(err)
	}
}
