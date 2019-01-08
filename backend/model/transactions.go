package model

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/iost-official/explorer/backend/model/db"
	"github.com/iost-official/explorer/backend/util"
)

/// this struct is used as json to return
type TxnDetail struct {
	Hash          string  `json:"txHash"`
	BlockNumber   int64   `json:"blockHeight"`
	From          string  `json:"from"`
	To            string  `json:"to"`
	Amount        float64 `json:"amount"`
	GasLimit      float64 `json:"gasLimit"`
	GasPrice      float64 `json:"price"`
	Age           string  `json:"age"`
	UTCTime       string  `json:"utcTime"`
	Code          string  `json:"code"`
	StatusCode    int32   `json:"statusCode"`
	StatusMessage string  `json:"statusMessage"`
	Contract      string  `json:"contract"`
	ActionName    string  `json:"actionName"`
	Data          string  `json:"data"`
}

type TxJson struct {
	Hash        string  `json:"hash"`
	BlockNumber int64   `json:"blockNumber"`
	From        string  `json:"from"`
	To          string  `json:"to"`
	Amount      float64 `json:"amount"`
	GasLimit    float64 `json:"gasLimit"`
	GasPrice    float64 `json:"gasPrice"`
	Age         string  `json:"age"`
	UTCTime     string  `json:"utcTime"`
}

func ConvertTxJsons(txs []*db.TxStore) []*TxJson {
	var ret []*TxJson
	for _, tx := range txs {
		txnOut := &TxJson{
			Hash:        tx.Tx.Hash,
			BlockNumber: tx.BlockNumber,
			From:        tx.Tx.Publisher,
			To:          tx.Tx.Actions[0].Contract,
			GasLimit:    tx.Tx.GasLimit,
			GasPrice:    tx.Tx.GasRatio,
			Age:         util.ModifyIntToTimeStr(tx.Tx.Time / 1e9),
			UTCTime:     util.FormatUTCTime(tx.Tx.Time),
		}

		if tx.Tx.Actions[0].Contract == "token.iost" && tx.Tx.Actions[0].ActionName == "transfer" {
			var params []string
			err := json.Unmarshal([]byte(tx.Tx.Actions[0].Data), &params)
			if err == nil && len(params) == 5 && params[0] == "iost" {
				txnOut.From = params[1]
				txnOut.To = params[2]
				f, _ := strconv.ParseFloat(params[3], 64)
				txnOut.Amount = f
			}
		}
		ret = append(ret, txnOut)
	}
	return ret
}

func ConvertTxsOutputs(tx []*db.TxStore) []*TxnDetail {
	var ret []*TxnDetail
	for _, t := range tx {
		ret = append(ret, ConvertTxOutput(t))
	}
	return ret
}

/// convert FlatTx to TxnDetail
func ConvertTxOutput(tx *db.TxStore) *TxnDetail {
	txnOut := &TxnDetail{
		Hash:          tx.Tx.Hash,
		BlockNumber:   tx.BlockNumber,
		From:          tx.Tx.Publisher,
		To:            tx.Tx.Actions[0].Contract,
		GasLimit:      tx.Tx.GasLimit,
		GasPrice:      tx.Tx.GasRatio,
		Age:           util.ModifyIntToTimeStr(tx.Tx.Time / 1e9),
		UTCTime:       util.FormatUTCTime(tx.Tx.Time),
		Code:          "",
		StatusCode:    int32(tx.Tx.TxReceipt.StatusCode),
		StatusMessage: tx.Tx.TxReceipt.Message,
		Contract:      tx.Tx.Actions[0].Contract,
		ActionName:    tx.Tx.Actions[0].ActionName,
		Data:          tx.Tx.Actions[0].Data,
	}

	if tx.Tx.Actions[0].Contract == "token.iost" && tx.Tx.Actions[0].ActionName == "transfer" {
		var params []string
		err := json.Unmarshal([]byte(tx.Tx.Actions[0].Data), &params)
		if err == nil && len(params) == 5 && params[0] == "iost" {
			txnOut.From = params[1]
			txnOut.To = params[2]
			f, _ := strconv.ParseFloat(params[3], 64)
			txnOut.Amount = f
		}
	}

	if tx.Tx.Actions[0].ActionName == "SetCode" {
		/*      c := new(contract.Contract) */
		// dataArr := tx.Action.Data

		// // remove comma if necessary
		// if dataArr[len(dataArr)-2] == ',' {
		// dataArr = dataArr[:len(dataArr)-2] + "]"
		// }

		// var code []string
		// json.Unmarshal([]byte(dataArr), &code)

		// c.B64Decode(code[0])
		/* txnOut.Code = c.Code */
	}

	return txnOut
}

func GetDetailTxn(txHash string) (*TxnDetail, error) {
	tx, err := db.GetTxByHash(txHash)

	if err != nil {
		log.Printf("transaction not found. txHash:%v, err=%v", txHash, err)
		return nil, err
	}

	txnOut := ConvertTxOutput(tx)

	return txnOut, nil
}

/// get a list of transactions for a specific page using account and block
func GetFlatTxnSlicePage(page, eachPageNum, block int64) ([]*TxnDetail, error) {
	lastPageNum, err := db.GetTxTotalPageCnt(eachPageNum, block)
	if err != nil {
		return nil, err
	}

	if lastPageNum == 0 {
		return []*TxnDetail{}, nil
	}

	if page > lastPageNum {
		page = lastPageNum
	}

	start := int((page - 1) * eachPageNum)
	txnsFlat, err := db.GetTxs(block, start, int(eachPageNum))

	if err != nil {
		return nil, err
	}

	return ConvertTxsOutputs(txnsFlat), nil
}
