package main

import (
	"fmt"
	"github.com/iost-official/explorer/backend/model/db"
	"github.com/iost-official/explorer/backend/task/cron"
	"github.com/spf13/viper"
	"sync"
)

var ws = new(sync.WaitGroup)

func main()  {
	viper.SetConfigName("config")
	viper.AddConfigPath("~/go/src/github.com/iost-official/explorer/backend")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	db.InitConfig()
	ws.Add(5)
	go cron.UpdateBlocks(ws)
	go cron.ProcessFailedSyncBlocks(ws)
	go cron.UpdateTxns(ws)
	go cron.UpdateBlockPay(ws)
	go cron.UpdateAccounts(ws)
	ws.Wait()
}
