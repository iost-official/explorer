package blkchain

import (
	"context"
	"time"

	"github.com/iost-official/explorer/backend/model/blkchain/rpc"
	"github.com/iost-official/explorer/backend/util/transport"
)

func GetCurrentBlockHeight() (int64, error) {
	conn, err := transport.GetGRPCClient(RPCAddress)
	if err != nil {
		return 0, err
	}

	client := rpc.NewApisClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rs, err := client.GetHeight(ctx, &rpc.VoidReq{})
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
