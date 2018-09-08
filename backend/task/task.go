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
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	db.InitConfig()

	// start tasks
	ws.Add(6)
	go cron.UpdateBlocks(ws)
	go cron.ProcessFailedSyncBlocks(ws)
	go cron.UpdateTxns(ws)
	go cron.UpdateBlockPay(ws)
	go cron.UpdateAccounts(ws)
	go cron.CheckRestart(ws)
	ws.Wait()
}
