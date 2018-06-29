package blockchain

import (
	"errors"
	"time"

	"github.com/iost-official/prototype/transport"

	"golang.org/x/net/context"

	"github.com/iost-official/prototype/rpc"
)

var ErrEmptyBlock = errors.New("no block found.")

func GetBlocks(start, limit int) ([]*rpc.BlockInfo, error) {
	// wait implement...
	return nil, nil
}

func GetBlockByLayer(layer int64) (*rpc.BlockInfo, error) {
	conn, err := transport.GetGRPCClient(RPCAddress)
	if err != nil {
		return nil, err
	}

	c := rpc.NewCliClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	bInfo, err := c.GetBlock(ctx, &rpc.BlockKey{Layer: layer})
	if err != nil {
		return nil, err
	}

	return bInfo, nil
}

func GetBlockByHeight(height int64) (*rpc.BlockInfo, error) {
	conn, err := transport.GetGRPCClient(RPCAddress)
	if err != nil {
		return nil, err
	}

	c := rpc.NewCliClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	bkey := &rpc.BlockKey{Layer: height}
	bInfo, err := c.GetBlockByHeight(ctx, bkey)
	if err != nil {
		return nil, err
	}

	return bInfo, nil
}

func GetTopBlock() (*rpc.BlockInfo, error) {
	blk, err := GetBlockByLayer(0)
	if err != nil {
		return nil, err
	}

	if blk == nil {
		return nil, ErrEmptyBlock
	}

	return blk, nil
}

func GetBlockLastPage(eachPage int64) int64 {
	var pageLast int64
	if topBlock, err := GetTopBlock(); err == nil {
		if topBlock.Head.Number % eachPage == 0 {
			pageLast = topBlock.Head.Number / eachPage
		} else {
			pageLast = topBlock.Head.Number / eachPage + 1
		}
	}

	return pageLast
}
