package blkchain

import "github.com/spf13/viper"

const (

	SystemContract     = "iost.system"
	SystemTransferFunc = "Transfer"
)

var (
	RPCAddress = "13.237.151.211:30002"
)

func InitConfig() {
	RPCAddress = viper.GetString("rpcHost")
}
