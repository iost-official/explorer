package blockchain

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	RPCAddress = "127.0.0.1:30002"
)

func InitConfig() {
	RPCAddress = viper.GetString("rpcHost")
	fmt.Println("RPCHost:", RPCAddress)
}
