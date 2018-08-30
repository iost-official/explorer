package blkchain

import (
	"encoding/json"

	"github.com/iost-official/Go-IOS-Protocol/crypto"
)

const (
	RPCAddress                 = "http://47.75.223.44:30301"
	BlockHeightRPCUrl          = RPCAddress + "/getHeight"
	GetTxByHashRPCUrlPrefix    = RPCAddress + "/getTxByHash/"
	GetBlockByHashRPCUrlPrefix = RPCAddress + "/getBlockByHash/"
	GetBlockByNumRPCUrlPrefix  = RPCAddress + "/getBlockByNum/"
	GetBalanceRPCUrlPrefix     = RPCAddress + "/getBalance/"
)

type BlockHeight struct {
	Height json.Number `json:"height,omitempty"`
}

type Balance struct {
	Balance json.Number `json:"balance,omitempty"`
}

type ActionRaw struct {
	Contract   string `json:"contract,omitempty"`
	ActionName string `json:"actionName,omitempty"`
	Data       string `json:"data,omitempty"`
}

type Tx struct {
	Time       json.Number            `json:"time,omitempty"`
	Expiration json.Number            `json:"expiration,omitempty"`
	GasLimit   json.Number            `json:"gasLimit,omitempty"`
	GasPrice   json.Number            `json:"gasPrice,omitempty"`
	Actions    []*ActionRaw           `json:"actions,omitempty"`
	Signers    [][]byte               `json:"signers,omitempty"`
	Signs      []*crypto.SignatureRaw `json:"signs,omitempty"`
	Publisher  *crypto.SignatureRaw   `json:"publisher,omitempty"`
}

type BlockHead struct {
	Version    json.Number `json:"version,omitempty"`
	ParentHash []byte      `json:"parentHash,omitempty"`
	TxsHash    []byte      `json:"txsHash,omitempty"`
	MerkleHash []byte      `json:"merkleHash,omitempty"`
	Info       []byte      `json:"info,omitempty"`
	Number     json.Number `json:"number,omitempty"`
	Witness    string      `json:"witness,omitempty"`
	Time       json.Number `json:"time,omitempty"`
}

type BlockInfo struct {
	Head   *BlockHead `json:"head,omitempty"`
	Txs    []*Tx      `json:"txs,omitempty"`
	Txhash []string   `json:"txhash,omitempty"`
}
