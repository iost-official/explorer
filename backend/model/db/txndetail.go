package db

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	"github.com/iost-official/prototype/common"
	"github.com/iost-official/prototype/rpc"
)

type MgoTx struct {
	MgoSourceId bson.ObjectId    `json:"mgo_source_id"`
	BlockHeight int64            `json:"block_height"`
	TxHash      []byte           `json:"tx_hash"`
	Time        int64            `json:"time"`
	Nonce       int64            `json:"nonce"`
	Publisher   common.Signature `json:"publisher"`
	GasLimit    int64            `json:"gas_limit"`
	Price       float64          `json:"price"`
	From        string           `json:"from"`
	To          string           `json:"to"`
	Amount      float64          `json:"amount"`
	Code        string           `json:"code"`
}

func GetTxnDetail(start, limit, block int, address string) ([]*MgoTx, error) {
	txnDC, err := GetCollection("txnsdetail")
	if err != nil {
		log.Println("UpdateTxns get txns collection error:", err)
		return nil, err
	}

	var pip []bson.M

	if address != "" {
		pip = []bson.M{
			bson.M{
				"$match": bson.M{
					"$or": []bson.M{
						bson.M{"from": address},
						bson.M{"to": address},
					},
				},
			},
		}
	}

	if block >= 0 {
		pip = []bson.M{
			bson.M{
				"$match": bson.M{
					"blockheight": block,
				},
			},
		}
	}

	pip = append(pip, []bson.M{
		bson.M{
			"$sort": bson.M{"blockheight": -1},
		},
		bson.M{
			"$skip": start,
		},
		bson.M{
			"$limit": limit,
		},
	}...)

	var txnsDetail []*MgoTx
	err = txnDC.Pipe(pip).All(&txnsDetail)
	if err != nil {
		return nil, err
	}

	return txnsDetail, nil
}

func GetTxnDetailListByHeight(height int64) ([]*MgoTx, error) {
	txnDC, err := GetCollection("txnsdetail")
	if err != nil {
		log.Println("UpdateTxns get txns collection error:", err)
		return nil, err
	}

	listQuery := bson.M{
		"blockheight": height,
	}

	var detailList []*MgoTx
	err = txnDC.Find(listQuery).All(&detailList)
	if err != nil {
		return nil, err
	}

	return detailList, nil
}

func GetTxnDetailByHash(txHash string) (*MgoTx, error) {
	txnDC, err := GetCollection("txnsdetail")
	if err != nil {
		log.Println("UpdateTxns get txns collection error:", err)
		return nil, err
	}

	txHashBytes := common.Base58Decode(txHash)
	if err != nil {
		return nil, err
	}

	query := bson.M{
		"txhash": txHashBytes,
	}
	var txn *MgoTx
	err = txnDC.Find(query).One(&txn)

	return txn, err
}

func GetTxnDetailByKey(tkey *rpc.TransactionKey) (*MgoTx, error) {
	txnDC, err := GetCollection("txnsdetail")
	if err != nil {
		log.Println("GetTxnDetailByKey get txns collection error:", err)
		return nil, err
	}

	txnQuery := bson.M{
		"nonce":            tkey.Nonce,
		"publisher.pubkey": tkey.Publisher,
	}

	var txnDetail *MgoTx
	err = txnDC.Find(txnQuery).One(&txnDetail)
	if err != nil {
		return nil, err
	}

	return txnDetail, err
}

func GetTotalTxnDetailLen(address string, block int64) (int, error) {
	txnDC, err := GetCollection("txnsdetail")
	if err != nil {
		log.Println("UpdateTxns get txns collection error:", err)
		return 0, err
	}

	var query bson.M
	if address != "" {
		query = bson.M{
			"$or": []bson.M{
				bson.M{"from": address},
				bson.M{"to": address},
			},
		}
	}
	if block >= 0 {
		query = bson.M{
			"blockheight": block,
		}
	}

	return txnDC.Find(query).Count()
}

func GetTxDetailLastPage(eachPage int64) (int64, error) {
	totalLen, err := GetTotalTxnDetailLen("", -1)
	if err != nil {
		return 0, err
	}
	txsInt64Len := int64(totalLen)

	var pageLast int64
	if txsInt64Len%eachPage == 0 {
		pageLast = txsInt64Len / eachPage
	} else {
		pageLast = txsInt64Len / eachPage + 1
	}

	return pageLast, nil
}

func GetTopTxnDetail() (*MgoTx, error) {
	txnDC, err := GetCollection("txnsdetail")
	if err != nil {
		log.Println("GetTopTxnDetail get txnsdetail collection error:", err)
		return nil, err
	}

	var (
		emptyQuery   interface{}
		topTxnDetail *MgoTx
	)
	err = txnDC.Find(emptyQuery).Sort("-_id").Limit(1).One(&topTxnDetail)
	if err != nil {
		log.Println("GetTopTxnDetail error:", err)
		return nil, err
	}

	return topTxnDetail, nil
}

func GetTxnListByAccount(account string, start, limit int) ([]*MgoTx, error) {
	txnDC, err := GetCollection("txnsdetail")
	if err != nil {
		log.Println("GetTxnListByAccount get txnsdetail collection error:", err)
		return nil, err
	}

	var txnList []*MgoTx
	pip := []bson.M{
		bson.M{
			"$match": bson.M{
				"$or": []bson.M{
					bson.M{"from": account},
					bson.M{"to": account},
				},
			},
		},
		bson.M{
			"$sort": bson.M{"blockheight": -1},
		},
		bson.M{
			"$skip": start,
		},
		bson.M{
			"$limit": limit,
		},
	}

	err = txnDC.Pipe(pip).All(&txnList)
	if err != nil {
		return nil, err
	}

	return txnList, nil
}

func GetTxnDetailLenByAccount(account string) (int, error) {
	txnDC, err := GetCollection("txnsdetail")
	if err != nil {
		log.Println("GetTopTxnDetail get txnsdetail collection error:", err)
		return 0, err
	}

	fromLen, err := txnDC.Find(bson.M{"from": account}).Count()
	if err != nil {
		log.Println("GetTxnDetailLenByAccount get from len error:", err)
	}
	toLen, err := txnDC.Find(bson.M{"to": account}).Count()
	if err != nil {
		log.Println("GetTxnDetailLenByAccount get to len error:", err)
	}
	return fromLen + toLen, err
}

func GetTxDetailLastPageWithAddress(eachPage int64, address string) (int64, error) {
	intLen, err := GetTxnDetailLenByAccount(address)
	if err != nil {
		return 0, err
	}

	txsInt64Len := int64(intLen)

	var pageLast int64
	if txsInt64Len%eachPage == 0 {
		pageLast = txsInt64Len / eachPage
	} else {
		pageLast = txsInt64Len / eachPage
	}

	return pageLast, nil
}

func GetTxDetailLastPageWithBlk(eachPage int64, blk int64) (int64, error) {
	txnDC, err := GetCollection("txnsdetail")
	if err != nil {
		log.Println("GetTxDetailLastPageWithAddress get collection error:", err)
		return 0, err
	}

	query := bson.M{
		"blockheight": blk,
	}

	intLen, err := txnDC.Find(query).Count()
	if err != nil {
		return 0, err
	}

	txsInt64Len := int64(intLen)

	var pageLast int64
	if txsInt64Len%eachPage == 0 {
		pageLast = txsInt64Len / eachPage
	} else {
		pageLast = txsInt64Len / eachPage
	}

	return pageLast, nil
}
