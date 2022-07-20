package blockchain

import (
	"context"
	"time"

	"github.com/GincoInc/iost-explorer/backend/model/blockchain/rpcpb"
	"github.com/GincoInc/iost-explorer/backend/util/transport"
)

func GetBlockByNum(num int64, complete bool) (*rpcpb.BlockResponse, error) {
	conn, err := transport.GetGRPCClient(RPCAddress)
	if err != nil {
		return nil, err
	}

	client := rpcpb.NewApiServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return client.GetBlockByNumber(ctx, &rpcpb.GetBlockByNumberRequest{
		Number:   num,
		Complete: complete,
	})
}

func GetBlockByHash(hash string, complete bool) (*rpcpb.BlockResponse, error) {
	conn, err := transport.GetGRPCClient(RPCAddress)
	if err != nil {
		return nil, err
	}

	client := rpcpb.NewApiServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return client.GetBlockByHash(ctx, &rpcpb.GetBlockByHashRequest{
		Hash:     hash,
		Complete: complete,
	})
}

func GetTxByHash(hash string) (*rpcpb.TransactionResponse, error) {
	conn, err := transport.GetGRPCClient(RPCAddress)
	if err != nil {
		return nil, err
	}

	client := rpcpb.NewApiServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return client.GetTxByHash(ctx, &rpcpb.TxHashRequest{
		Hash: hash,
	})
}

func GetTxReceiptByTxHash(hash string) (*rpcpb.TxReceipt, error) {
	conn, err := transport.GetGRPCClient(RPCAddress)
	if err != nil {
		return nil, err
	}

	client := rpcpb.NewApiServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return client.GetTxReceiptByTxHash(ctx, &rpcpb.TxHashRequest{
		Hash: hash,
	})
}

func GetAccount(name string, byLongestChain bool) (*rpcpb.Account, error) {
	conn, err := transport.GetGRPCClient(RPCAddress)
	if err != nil {
		return nil, err
	}

	client := rpcpb.NewApiServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return client.GetAccount(ctx, &rpcpb.GetAccountRequest{
		Name:           name,
		ByLongestChain: byLongestChain,
	})
}

func GetContract(id string, byLongestChain bool) (*rpcpb.Contract, error) {
	conn, err := transport.GetGRPCClient(RPCAddress)
	if err != nil {
		return nil, err
	}
	client := rpcpb.NewApiServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return client.GetContract(ctx, &rpcpb.GetContractRequest{
		Id:             id,
		ByLongestChain: byLongestChain,
	})
}

func GetTokenBalance(account, token string, byLongestChain bool) (*rpcpb.GetTokenBalanceResponse, error) {
	conn, err := transport.GetGRPCClient(RPCAddress)
	if err != nil {
		return nil, err
	}

	client := rpcpb.NewApiServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return client.GetTokenBalance(ctx, &rpcpb.GetTokenBalanceRequest{
		Account:        account,
		Token:          token,
		ByLongestChain: byLongestChain,
	})
}
