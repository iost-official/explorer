package config

import (
	"fmt"

	"github.com/iost-official/iost-api/model/blockchain"
	"github.com/iost-official/iost-api/model/db"
	"github.com/spf13/viper"
)

func ReadConfig(configPath string) {
	viper.SetConfigName("config")
	if configPath == "" {
		viper.AddConfigPath("$GOPATH/src/github.com/iost-official/iost-api/config")
	} else {
		viper.AddConfigPath(configPath)
	}
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	db.InitConfig()
	blockchain.InitConfig()
}
