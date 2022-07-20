package blockchain

import (
	"fmt"
	"github.com/GincoInc/iost-explorer/backend/util"
	"github.com/spf13/viper"
)

var (
	RPCAddress  = "127.0.0.1:30002"
	PasswordKey = []byte("temp")
	BPRegisterPassword = "temp"
)

func InitConfig() {
	RPCAddress = viper.GetString("rpcHost")
	BPRegisterPassword = viper.GetString("BPRegisterPassword")
	fmt.Println("RPCHost:", RPCAddress)
	bytePassword, err := util.ReadPassword("Enter Password:  ")
	if err != nil {
		fmt.Println(err)
	}
	PasswordKey = bytePassword
}
