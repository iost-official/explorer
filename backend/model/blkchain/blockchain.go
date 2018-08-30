package blkchain

import (
	"encoding/json"
	"fmt"
)

func GetCurrentBlockHeight() (int64, error) {
	body, err := RPCRequest("GET", BlockHeightRPCUrl, nil)
	if err != nil {
		return 0, err
	}

	var blkHeight *BlockHeight
	err = json.Unmarshal(body, &blkHeight)
	if err != nil {
		return 0, err
	}

	return blkHeight.Height.Int64()
}

func GetTxByHash(hash string) (*Tx, error) {
	getTxByHashRPCUrl := GetTxByHashRPCUrlPrefix + hash
	body, err := RPCRequest("GET", getTxByHashRPCUrl, nil)
	if err != nil {
		return nil, err
	}

	var tx *Tx
	err = json.Unmarshal(body, &tx)

	return tx, err
}

func GetBlockByHash(hash string, complete bool) (*BlockInfo, error) {
	getBlockByHashRPCUrl := fmt.Sprintf("%s%s/%t", GetBlockByHashRPCUrlPrefix, hash, complete)
	body, err := RPCRequest("GET", getBlockByHashRPCUrl, nil)
	if err != nil {
		return nil, err
	}

	var block *BlockInfo
	err = json.Unmarshal(body, &block)

	return block, err
}

func GetBlockByNum(num int64, complete bool) (*BlockInfo, error) {
	getBlockByNumRPCUrl := fmt.Sprintf("%s%d/%t", GetBlockByNumRPCUrlPrefix, num, complete)
	body, err := RPCRequest("GET", getBlockByNumRPCUrl, nil)
	if err != nil {
		return nil, err
	}

	var block *BlockInfo
	err = json.Unmarshal(body, &block)

	return block, err
}

func GetBalance(address string) (int64, error) {
	getBalanceRPCUrl := GetBalanceRPCUrlPrefix + address + "/0"
	body, err := RPCRequest("GET", getBalanceRPCUrl, nil)
	if err != nil {
		return 0, err
	}

	println(string(body))
	var balance *Balance
	err = json.Unmarshal(body, &balance)
	if err != nil {
		return 0, err
	}

	return balance.Balance.Int64()
}
