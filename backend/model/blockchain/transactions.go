package blockchain

import (
	"time"

	"golang.org/x/net/context"

	"github.com/iost-official/prototype/core/tx"
	"github.com/iost-official/prototype/rpc"
	"github.com/iost-official/prototype/transport"
)

func GetTxnByKey(tkey *rpc.TransactionKey) (*tx.Tx, error) {
	conn, err := transport.GetGRPCClient(RPCAddress)
	if err != nil {
		return nil, err
	}

	c := rpc.NewCliClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	t, err := c.GetTransaction(ctx, tkey)
	if err != nil {
		return nil, err
	}

	trans := new(tx.Tx)
	err = trans.Decode(t.Tx)
	return trans, err
}

func GetTxnByHash(txHash []byte) (*tx.Tx, error) {
	conn, err := transport.GetGRPCClient(RPCAddress)
	if err != nil {
		return nil, err
	}

	c := rpc.NewCliClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	th := &rpc.TransactionHash{
		Hash: txHash,
	}
	t, err := c.GetTransactionByHash(ctx, th)
	if err != nil {
		return nil, err
	}

	trans := new(tx.Tx)
	err = trans.Decode(t.Tx)
	return trans, err
}
