package model

import (
	"explorer/model/db"

	"github.com/iost-official/prototype/rpc"
	"explorer/model/blockchain"
	"github.com/iost-official/prototype/common"
)

const (
	TxHashShortLen = 19
	FromToShortLen = 25
)

type TransactionOutput struct {
	TxHash      string  `json:"tx_hash"`
	BlockHeight int64   `json:"block_height"`
	From        string  `json:"from"`
	To          string  `json:"to"`
	Amount      float64 `json:"amount"`
	GasLimit    int64   `json:"gas_limit"`
	Price       float64 `json:"price"`
	Age         string  `json:"age"`
	UTCTime     string  `json:"utc_time"`
	Code        string  `json:"code"`
}

func GetTransaction(page, eachPageNum, block int64, address string) ([]*TransactionOutput, error) {
	lastPage, err := db.GetTxDetailLastPage(eachPageNum)
	if err != nil {
		return nil, err
	}

	if page > lastPage {
		page = lastPage
	}

	start := int((page - 1) * eachPageNum)
	txns, err := db.GetTxnDetail(start, int(eachPageNum), int(block), address)
	if err != nil {
		return nil, err
	}

	var txnOutputList []*TransactionOutput
	for _, v := range txns {
		txnOutputList = append(txnOutputList, GenerateTxnOutput(v))
	}

	return txnOutputList, nil
}

func GetTxnByKey(blkHeight int64, tkey *rpc.TransactionKey) (*TransactionOutput, error) {
	trans, err := db.GetTxnDetailByKey(tkey)
	if err != nil {
		return nil, err
	}

	return GenerateTxnOutput(trans), nil
}

func GenerateTxnOutput(trans *db.MgoTx) *TransactionOutput {
	// nano secs to secs
	timestamp := trans.Time / 1000000000

	//txHash := hex.EncodeToString(trans.TxHash)
	txHash := common.Base58Encode(trans.TxHash)
	output := &TransactionOutput{
		TxHash:      txHash,
		BlockHeight: trans.BlockHeight,
		From:        trans.From,
		To:          trans.To,
		Amount:      trans.Amount,
		GasLimit:    trans.GasLimit,
		Price:       trans.Price,
		Age:         modifyIntToTimeStr(timestamp),
		UTCTime:     formatUTCTime(timestamp),
		Code:        trans.Code,
	}

	if output.To == "Bet" {
		output.To = blockchain.BetHash
	}

	return output
}
