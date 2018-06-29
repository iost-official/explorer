package model

import (
	"encoding/hex"

	"explorer/model/db"

	"github.com/iost-official/prototype/rpc"
)

type BlockOutput struct {
	Height        int64                 `json:"height"`
	ParentHash    string                `json:"parent_hash"`
	BlockHash     string                `json:"block_hash"`
	Signature     string                `json:"signature"`
	Witness       string                `json:"witness"`
	Age           string                `json:"age"`
	UTCTime       string                `json:"utc_time"`
	Timestamp     int64                 `json:"timestamp"`
	Txn           int64                 `json:"txn"`
	TxList        []*rpc.TransactionKey `json:"tx_list"`
	TotalGasLimit int64                 `json:"total_gas_limit"`
	AvgGasPrice   float64               `json:"avg_gas_price"`
}

func GetBlock(page, eachPageNum int64) ([]*BlockOutput, error) {
	start := int((page - 1) * eachPageNum)

	blkInfoList, err := db.GetBlocks(start, int(eachPageNum))
	if err != nil {
		return nil, err
	}

	var blkHeightList []int64

	for _, v := range blkInfoList {
		blkHeightList = append(blkHeightList, v.Head.Number)
	}

	payMap, _ := db.GetBlockPayListByHeight(blkHeightList)

	var blockOutputList []*BlockOutput
	for _, v := range blkInfoList {
		output := GenerateBlockOutput(v)
		if pay, ok := payMap[v.Head.Number]; ok {
			output.TotalGasLimit = pay.TotalGasLimit
			output.AvgGasPrice = pay.AvgGasPrice
		}
		blockOutputList = append(blockOutputList, output)
	}

	return blockOutputList, nil
}

func GenerateBlockOutput(bInfo *rpc.BlockInfo) *BlockOutput {
	timestamp := ConvertSlotTimeToTimeStamp(bInfo.Head.Time)
	return &BlockOutput{
		Height:     bInfo.Head.Number,
		ParentHash: hex.EncodeToString(bInfo.Head.ParentHash),
		BlockHash:  hex.EncodeToString(bInfo.Head.BlockHash),
		Signature:  hex.EncodeToString(bInfo.Head.Signature),
		Witness:    bInfo.Head.Witness,
		Age:        modifyBlockIntToTimeStr(timestamp),
		UTCTime:    formatUTCTime(timestamp),
		Timestamp:  timestamp,
		Txn:        bInfo.Txcnt,
		TxList:     bInfo.TxList,
	}
}
