package blkchain

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/iost-official/explorer/backend/model/blkchain/rpc"
	"github.com/iost-official/explorer/backend/util/transport"
	"github.com/iost-official/go-iost/account"
	"github.com/iost-official/go-iost/common"
	"github.com/iost-official/go-iost/core/tx"
	"github.com/iost-official/go-iost/crypto"
)

func GetCurrentBlockHeight() (int64, error) {
	conn, err := transport.GetGRPCClient(RPCAddress)
	if err != nil {
		return 0, err
	}

	client := rpc.NewApisClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rs, err := client.GetHeight(ctx, &empty.Empty{})
	if err != nil {
		return 0, err
	}

	return rs.Height, nil
}

func GetTxByHash(hash string) (*rpc.TxRes, error) {
	conn, err := transport.GetGRPCClient(RPCAddress)
	if err != nil {
		return nil, err
	}

	client := rpc.NewApisClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rs, err := client.GetTxByHash(ctx, &rpc.HashReq{
		Hash: hash,
	})
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func GetTxReceiptByHash(hash string) (*rpc.TxReceiptRes, error) {
	conn, err := transport.GetGRPCClient(RPCAddress)
	if err != nil {
		return nil, err
	}

	client := rpc.NewApisClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rs, err := client.GetTxReceiptByHash(ctx, &rpc.HashReq{
		Hash: hash,
	})
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func GetTxReceiptByTxHash(hash string) (*rpc.TxReceiptRes, error) {
	conn, err := transport.GetGRPCClient(RPCAddress)
	if err != nil {
		return nil, err
	}

	client := rpc.NewApisClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rs, err := client.GetTxReceiptByTxHash(ctx, &rpc.HashReq{
		Hash: hash,
	})
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func GetBlockByHash(hash string, complete bool) (*rpc.BlockInfo, error) {
	conn, err := transport.GetGRPCClient(RPCAddress)
	if err != nil {
		return nil, err
	}

	client := rpc.NewApisClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return client.GetBlockByHash(ctx, &rpc.BlockByHashReq{
		Hash:     hash,
		Complete: complete,
	})
}

func GetBlockByNum(num int64, complete bool) (*rpc.BlockInfo, error) {
	conn, err := transport.GetGRPCClient(RPCAddress)
	if err != nil {
		return nil, err
	}

	client := rpc.NewApisClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return client.GetBlockByNum(ctx, &rpc.BlockByNumReq{
		Num:      num,
		Complete: complete,
	})
}

func GetBalance(address string) (int64, error) {
	conn, err := transport.GetGRPCClient(RPCAddress)
	if err != nil {
		return 0, err
	}

	client := rpc.NewApisClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rs, err := client.GetBalance(ctx, &rpc.GetBalanceReq{
		ID:              address,
		UseLongestChain: false,
	})
	if err != nil {
		return 0, err
	}

	return rs.Balance, nil
}

/*
 * run example:
 *
 *hash, err := Transfer("IOST6Jymdka3EFLAv8954MJ1nBHytNMwBkZfcXevE2PixZHsSrRkbR",
 * 		"pwd",
 *		1233,
 *		100000,
 *		1,
 *		100,
 *		"2Hoo4NAoFsx9oat6qWawHtzqFYcA3VS7BLxPowvKHFPM")
 */
func Transfer(from, to string, amount, gasLimit, gasPrice, expiration int64, privKey string) ([]byte, error) {
	transferData := fmt.Sprintf(`["%s", "%s", %d]`, from, to, amount)
	action := tx.NewAction(SystemContract, SystemTransferFunc, transferData)
	actions := []*tx.Action{&action}

	trx := tx.NewTx(actions, [][]byte{}, gasLimit, gasPrice, time.Now().Add(time.Second*time.Duration(expiration)).UnixNano())

	acc, err := account.NewAccount(common.Base58Decode(privKey), crypto.Ed25519)
	if err != nil {
		return nil, err
	}
	stx, err := tx.SignTx(trx, acc)

	conn, err := grpc.Dial(RPCAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := rpc.NewApisClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := client.SendRawTx(ctx, &rpc.RawTxReq{
		Data: stx.Encode(),
	})
	if err != nil {
		return nil, err
	}

	return []byte(resp.Hash), nil
}
