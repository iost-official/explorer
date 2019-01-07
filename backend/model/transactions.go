package model

import (
	"encoding/json"
	"log"

	"github.com/globalsign/mgo/bson"
	"github.com/iost-official/explorer/backend/model/db"
	"github.com/iost-official/explorer/backend/util"
	"github.com/iost-official/go-iost/core/contract"
)

/// this struct is used as json to return
type TxnDetail struct {
	Hash          string  `json:"txHash"`
	BlockNumber   int64   `json:"blockHeight"`
	From          string  `json:"from"`
	To            string  `json:"to"`
	Amount        float64 `json:"amount"`
	GasLimit      int64   `json:"gasLimit"`
	GasPrice      int64   `json:"price"`
	Age           string  `json:"age"`
	UTCTime       string  `json:"utcTime"`
	Code          string  `json:"code"`
	StatusCode    int32   `json:"statusCode"`
	StatusMessage string  `json:"statusMessage"`
	Contract      string  `json:"contract"`
	ActionName    string  `json:"actionName"`
	Data          string  `json:"data"`
}

func GetDetailTxn(txHash string) (TxnDetail, error) {
	txnC, err := db.GetCollection(db.CollectionFlatTx)

	if err != nil {
		log.Println("failed To open collection collectionTxs")
		return TxnDetail{}, err
	}

	var tx db.FlatTx

	err = txnC.Find(bson.M{"hash": txHash}).One(&tx)

	if err != nil {
		log.Println("transaction not found")
		return TxnDetail{}, err
	}

	txnOut := ConvertFlatTx2TxnDetail(&tx)

	return txnOut, nil
}

/// convert FlatTx to TxnDetail
func ConvertFlatTx2TxnDetail(tx *db.FlatTx) TxnDetail {
	txnOut := TxnDetail{
		Hash:        tx.Hash,
		BlockNumber: tx.BlockNumber,
		From:        tx.From,
		To:          tx.To,
		Amount:      tx.Amount,
		GasLimit:    tx.GasLimit,
		GasPrice:    tx.GasPrice,
		Contract:    tx.Action.Contract,
		ActionName:  tx.Action.ActionName,
		Data:        tx.Action.Data,
	}

	if tx.Action.ActionName == "SetCode" {
		c := new(contract.Contract)
		dataArr := tx.Action.Data

		// remove comma if necessary
		if dataArr[len(dataArr)-2] == ',' {
			dataArr = dataArr[:len(dataArr)-2] + "]"
		}

		var code []string
		json.Unmarshal([]byte(dataArr), &code)

		c.B64Decode(code[0])
		txnOut.Code = c.Code
	}

	txnOut.Age = util.ModifyIntToTimeStr(tx.Time / (1000 * 1000 * 1000))
	txnOut.UTCTime = util.FormatUTCTime(tx.Time / (1000 * 1000 * 1000))
	txnOut.StatusCode = tx.Receipt.StatusCode
	txnOut.StatusMessage = tx.Receipt.StatusMessage

	return txnOut
}

/// get a list of transactions for a specific page using account and block
func GetFlatTxnSlicePage(page, eachPageNum, block int64, address string) ([]*TxnDetail, error) {
	lastPageNum, err := db.GetFlatTxTotalPageCnt(eachPageNum, address, block)

	if lastPageNum == 0 {
		return []*TxnDetail{}, nil
	}

	if err != nil {
		return nil, err
	}

	if page > lastPageNum {
		page = lastPageNum
	}

	start := int((page - 1) * eachPageNum)
	txnsFlat, err := db.GetFlatTxnSlice(start, int(eachPageNum), int(block), address)

	if err != nil {
		return nil, err
	}

	var txnDetailList []*TxnDetail

	for _, v := range txnsFlat {
		td := ConvertFlatTx2TxnDetail(v)
		txnDetailList = append(txnDetailList, &td)
	}

	return txnDetailList, nil
}
