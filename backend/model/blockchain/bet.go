package blockchain

import (
	"fmt"
	"log"
	"time"

	"github.com/iost-official/prototype/transport"
	"explorer/model/db"

	"golang.org/x/net/context"

	"github.com/iost-official/prototype/rpc"
	"github.com/iost-official/prototype/core/state"
	"strconv"
	"github.com/iost-official/prototype/common"
	"explorer/model/cache"
)

const (
	BetHash = "9M8wiAvqkt6T8dtPaegjwfogbyogmMokJ6vbqW5wwaMJ"
	SendBetContract        = `--- main 合约主入口
-- LuckyBet
-- @gas_limit 100000000
-- @gas_price 0
-- @param_cnt 0
-- @return_cnt 0
function main()
	ok, r = Call("` + BetHash + `", "Bet", "%s", %d, %d, %d)
	Log(string.format("bet %%s", tostring(ok)))
	Log(string.format("bet r = %%s", tostring(r)))
	Assert(ok)
	Assert(r == 0)
end--f`
)

func init()  {
	go GetBetMainCode()
}

func SendBet(address, privKey string, luckyNumber, betAmount int) ([]byte, int, error) {
	conn, err := transport.GetGRPCClient(RPCAddress)
	if err != nil {
		return nil, 0, err
	}

	nonce, err := db.GetAddressNonce(address)
	if err != nil && err.Error() != "not found" {
		log.Println("SendBet GetAddressNonce error:", err)
	}

	contract := fmt.Sprintf(SendBetContract, address, luckyNumber, betAmount, nonce)

	transInfo := &rpc.TransInfo{
		Seckey: privKey,
		Nonce: nonce,
		Contract: contract,
	}
	log.Println(transInfo)
	err = db.IncAddressNonce(address)
	log.Println("SendBet IncAddressNonce error:", err)

	c := rpc.NewCliClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ret, err := c.Transfer(ctx, transInfo)
	if err != nil {
		return nil, 0, err
	}

	if ret.Code == 0 {
		return ret.Hash, int(nonce), nil
	}
	return nil, 0, nil
}

func GetTotalRounds(txHash string) (float64, error) {
	conn, err := transport.GetGRPCClient(RPCAddress)
	if err != nil {
		return 0, err
	}

	c := rpc.NewCliClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	key := &rpc.Key{
		S:txHash + "round",
	}

	st, err := c.GetState(ctx, key)
	if err != nil {
		return 0, err
	}

	v, err := state.ParseValue(st.Sv)
	if err != nil {
		return 0, err
	}

	if vFloat, ok := v.(*state.VFloat); ok {
		return vFloat.ToFloat64(), nil
	}

	return 0, nil
}

func GetRoundWithNumber(txHash string, number int) (string, error) {
	conn, err := transport.GetGRPCClient(RPCAddress)
	if err != nil {
		return "", err
	}

	c := rpc.NewCliClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	key := &rpc.Key{
		S: txHash + "round" + strconv.Itoa(number),
	}
	st, err := c.GetState(ctx, key)

	if err != nil {
		return "", err
	}

	return st.Sv, nil
}

func GetBetMainCode() {
	txn, _ := GetTxnByHash(common.Base58Decode(BetHash))

	if txn != nil && txn.Contract != nil && txn.Contract.Code() != "" {
		cache.GlobalCache.Set("betMainCode", txn, -1)
	}

	ticker := time.NewTicker(time.Minute)
	for _ = range ticker.C {
		txn, err := GetTxnByHash(common.Base58Decode(BetHash))
		if err != nil {
			continue
		}

		if txn.Contract != nil && txn.Contract.Code() != "" {
			cache.GlobalCache.Set("betMainCode", txn, -1)
		}
	}
}
