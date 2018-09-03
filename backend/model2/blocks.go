package model2

import (
	"github.com/iost-official/explorer/backend/model2/db"
	"log"
)

type BlockOutput struct {
	Height        int64    `json:"height"`
	ParentHash    string   `json:"parent_hash"`
	BlockHash     string   `json:"block_hash"`
	Witness       string   `json:"witness"`
	Age           string   `json:"age"`
	UTCTime       string   `json:"utc_time"`
	Timestamp     int64    `json:"timestamp"`
	TxList        []string `json:"tx_list"`
	TotalGasLimit int64    `json:"total_gas_limit"`
	AvgGasPrice   float64  `json:"avg_gas_price"`
}

func GetBlock(page, eachPageNum int64) ([]*BlockOutput, error) {
	start := int((page - 1) * eachPageNum)

	blkInfoList, err := db.GetBlocks(start, int(eachPageNum))
	if err != nil {
		return nil, err
	}

	var blkHeightList []int64

	for _, v := range blkInfoList {
		blkHeightList = append(blkHeightList, v.BlockNumber)
	}

	payMap, _ := db.GetBlockPayListByHeight(blkHeightList)

	var blockOutputList []*BlockOutput
	for _, v := range blkInfoList {
		output := GenerateBlockOutput(v)
		if pay, ok := payMap[v.BlockNumber]; ok {
			output.TotalGasLimit = pay.TotalGasLimit
			output.AvgGasPrice = pay.AvgGasPrice
		}

		txList, err := db.GetBlockTxnHashes(v.BlockNumber)

		if nil == err {
			output.TxList = *txList
		} else {
			log.Println("get block txn list fail", err)
		}

		blockOutputList = append(blockOutputList, output)
	}

	return blockOutputList, nil
}

func GenerateBlockOutput(bInfo *db.Block) *BlockOutput {
	//todo when rpc fix this, change it to normal
	timestamp := bInfo.Time * 3
	return &BlockOutput{
		Height:     bInfo.BlockNumber,
		ParentHash: bInfo.ParentHash,
		BlockHash:  bInfo.Hash,
		Witness:    bInfo.Witness,
		Age:        modifyBlockIntToTimeStr(timestamp),
		UTCTime:    formatUTCTime(timestamp),
		Timestamp:  timestamp,
	}
}
