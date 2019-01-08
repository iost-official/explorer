package model

import (
	"log"

	"github.com/iost-official/explorer/backend/model/blockchain/rpcpb"
	"github.com/iost-official/explorer/backend/model/db"
	"github.com/iost-official/explorer/backend/util"
)

type BlockOutput struct {
	Height        int64    `json:"height"`
	ParentHash    string   `json:"parentHash"`
	BlockHash     string   `json:"blockHash"`
	Witness       string   `json:"witness"`
	Age           string   `json:"age"`
	UTCTime       string   `json:"utcTime"`
	Timestamp     int64    `json:"timestamp"`
	TxList        []string `json:"txList"`
	Txn           int64    `json:"txn"`
	TotalGasLimit int64    `json:"totalGasLimit"`
	AvgGasPrice   float64  `json:"avgGasPrice"`
}

func GetBlock(page, eachPageNum int64) ([]*BlockOutput, error) {
	start := int((page - 1) * eachPageNum)

	blkInfoList, err := db.GetBlocks(start, int(eachPageNum))
	if err != nil {
		return nil, err
	}
	/*
		var blkHeightList []int64

		for _, v := range blkInfoList {
			blkHeightList = append(blkHeightList, v.Number)
		}

		payMap, _ := db.GetBlockPayListByHeight(blkHeightList)
	*/
	var blockOutputList []*BlockOutput
	for _, v := range blkInfoList {
		output := GenerateBlockOutput(v)
		/*      if pay, ok := payMap[v.BlockNumber]; ok { */
		// output.TotalGasLimit = pay.TotalGasLimit
		// output.AvgGasPrice = pay.AvgGasPrice
		/* } */

		blockOutputList = append(blockOutputList, output)
	}

	return blockOutputList, nil
}

func GenerateBlockOutput(bInfo *rpcpb.Block) *BlockOutput {
	//todo when rpc fix this, change it to normal
	txCount, err := db.GetTxCountByNumber(bInfo.Number)
	if err != nil {
		log.Printf("get txcount failed. number:%d, err=%v", bInfo.Number, err)
	}
	return &BlockOutput{
		Height:     bInfo.Number,
		ParentHash: bInfo.ParentHash,
		BlockHash:  bInfo.Hash,
		Witness:    bInfo.Witness,
		Txn:        int64(txCount),
		Age:        util.ModifyBlockIntToTimeStr(bInfo.Time),
		UTCTime:    util.FormatUTCTime(bInfo.Time),
		Timestamp:  bInfo.Time,
	}
}
