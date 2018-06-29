package blockchain

import (
	"fmt"
	"time"

	"github.com/iost-official/prototype/transport"

	"golang.org/x/net/context"

	"github.com/iost-official/prototype/core/state"
	"github.com/iost-official/prototype/rpc"
	"explorer/model/db"
	"log"
)

const (
	TransferIOSTOrigPrivKey = "BRpwCKmVJiTTrPFi6igcSgvuzSiySd7Exxj7LGfqieW9"
	TransferIOSTContract = `--- main 合约主入口
-- server1转账server2
-- @gas_limit 10000
-- @gas_price 0.001
-- @param_cnt 0
-- @return_cnt 0
function main()
	Transfer("2BibFrAhc57FAd3sDJFbPqjwskBJb5zPDtecPWVRJ1jxT","%s",%f)
end--f`
)

func GetBalanceByKey(key string) (float64, error) {
	conn, err := transport.GetGRPCClient(RPCAddress)
	if err != nil {
		return 0, err
	}

	c := rpc.NewCliClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	bInfo, err := c.GetBalance(ctx, &rpc.Key{S: key})
	if err != nil {
		return 0, err
	}

	v, err := state.ParseValue(bInfo.Sv)
	if err != nil {
		return 0, err
	}

	if vFloat, ok := v.(*state.VFloat); ok {
		return vFloat.ToFloat64(), nil
	}

	return 0, nil
}

func TransferIOSTToAddress(address string, amount float64) ([]byte, error) {
	conn, err := transport.GetGRPCClient(RPCAddress)
	if err != nil {
		return nil, err
	}

	contract := fmt.Sprintf(TransferIOSTContract, address, amount)

	nonce, err := db.GetAddressNonce(address)
	if err != nil && err.Error() != "not found" {
		log.Println("TransferIOSTToAddress GetAddressNonce error:", err)
	}

	transInfo := &rpc.TransInfo{
		Seckey: TransferIOSTOrigPrivKey,
		Nonce: nonce,
		Contract: contract,
	}
	err = db.IncAddressNonce(address)
	log.Println("TransferIOSTToAddress IncAddressNonce error:", err)

	c := rpc.NewCliClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ret, err := c.Transfer(ctx, transInfo)
	if err != nil {
		return nil, err
	}

	if ret.Code == 0 {
		return ret.Hash, nil
	}
	return nil, nil
}
