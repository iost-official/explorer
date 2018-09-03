package model2

import (
	"github.com/globalsign/mgo/bson"
	"github.com/iost-official/explorer/backend/model2/db"
	"log"
)

type TxnDetail struct {
	Hash        string  `json:"tx_hash"`
	BlockNumber int64   `json:"block_height"`
	From        string  `json:"from"`
	To          string  `json:"to"`
	Amount      float64 `json:"amount"`
	GasLimit    int64   `json:"gas_limit"`
	GasPrice    int64   `json:"price"`
	Age         string  `json:"age"`
	UTCTime     string  `json:"utc_time"`
	Code        string  `json:"code"`
}

func GetDetailTxn(txHash string) TxnDetail {
	txnC, err := db.GetCollection(db.CollectionFlatTx)

	if err != nil {
		log.Println("failed To open collection collectionTxs")
		return TxnDetail{}
	}

	var tx db.FlatTx

	err = txnC.Find(bson.M{"hash": txHash}).One(&tx)

	if err != nil {
		log.Println("transaction not found")
		return TxnDetail{}
	}

	txnOut := TxnDetail{
		Hash:        txHash,
		BlockNumber: tx.BlockNumber,
		From:        tx.From,
		To:          tx.To,
		Amount:      tx.Amount,
		GasLimit:    tx.GasLimit,
		GasPrice:    tx.GasPrice,
	}

	if tx.ActionName == "SetCode" {
		txnOut.Code = tx.Action.Data
	}

	txnOut.Age = modifyIntToTimeStr(tx.Time / (1000 * 1000 * 1000))
	txnOut.UTCTime = formatUTCTime(tx.Time / (1000 * 1000 * 1000))

	return txnOut
}
