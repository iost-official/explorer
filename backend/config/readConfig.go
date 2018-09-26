package config

import (
	"fmt"
	"github.com/iost-official/explorer/backend/model/blkchain"
	"github.com/iost-official/explorer/backend/model/db"
	"github.com/spf13/viper"
)

func ReadConfig () {
	viper.SetConfigName("config")
	viper.AddConfigPath("$GOPATH/src/github.com/iost-official/explorer/backend/config")
	//viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	db.InitConfig()
	db.EnsureCapped()
	blkchain.InitConfig()
}
