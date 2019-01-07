package db

import (
	"github.com/globalsign/mgo/bson"
)

/// get at most `limit` flat txns from start using block and account address
func GetFlatTxnSlice(start, limit, block int, address string) ([]*FlatTx, error) {
	txnDC := GetCollection(CollectionFlatTx)

	var pip []bson.M

	if address != "" {
		pip = []bson.M{
			bson.M{
				"$match": bson.M{
					"$or": []bson.M{
						bson.M{"from": address},
						bson.M{"to": address},
						bson.M{"publisher": address},
					},
				},
			},
		}
	}

	if block >= 0 {
		pip = []bson.M{
			bson.M{
				"$match": bson.M{
					"blockNumber": block,
				},
			},
		}
	}

	pip = append(pip, []bson.M{
		bson.M{
			"$sort": bson.M{"blockNumber": -1},
		},
		bson.M{
			"$skip": start,
		},
		bson.M{
			"$limit": limit,
		},
	}...)

	var flatTx []*FlatTx

	err := txnDC.Pipe(pip).All(&flatTx)

	if err != nil {
		return nil, err
	}

	return flatTx, nil
}

/// get length of transaction list using account and block number
func GetTotalFlatTxnLen(address string, block int64) (int, error) {
	txnDC := GetCollection(CollectionFlatTx)

	var query bson.M

	if address != "" {
		query = bson.M{
			"$or": []bson.M{
				bson.M{"from": address},
				bson.M{"to": address},
				bson.M{"publisher": address},
			},
		}
		return 0, nil
	} else if block >= 0 {
		query = bson.M{
			"blockNumber": block,
		}
	}

	return txnDC.Find(query).Count()
}

func GetFlatTxTotalPageCnt(eachPage int64, account string, block int64) (int64, error) {
	totalLen, err := GetTotalFlatTxnLen(account, block)

	if err != nil {
		return 0, err
	}

	txsInt64Len := int64(totalLen)
	pageMax := txsInt64Len / eachPage

	if txsInt64Len%eachPage != 0 {
		pageMax++
	}

	return pageMax, nil
}

/* func GetFlatTxPageCntWithAddress(eachPage int64, account string) (int64, error) { */
// intLen, err := GetFlatTxnLenByAccount(account)
// if err != nil {
// return 0, err
// }

// txsInt64Len := int64(intLen)

// var pageMax = txsInt64Len / eachPage

// if txsInt64Len%eachPage != 0 {
// pageMax++
// }

// return pageMax, nil
/* } */

func GetFlatTxPageCntWithBlk(eachPage int64, blk int64) (int64, error) {
	txnDC := GetCollection(CollectionFlatTx)

	query := bson.M{
		"blockNumber": blk,
	}

	intLen, err := txnDC.Find(query).Count()

	if err != nil {
		return 0, err
	}

	txsInt64Len := int64(intLen)

	var pageLast = txsInt64Len / eachPage

	if txsInt64Len%eachPage != 0 {
		pageLast++
	}

	return pageLast, nil
}

func GetTxnDetailByHash(txHash string) (*Tx, error) {
	txnDC := GetCollection(CollectionTxs)

	query := bson.M{
		"hash": txHash,
	}
	var txn *Tx
	err := txnDC.Find(query).One(&txn)

	return txn, err
}

func GetFlatTxnDetailByHash(txHash string) (*FlatTx, error) {
	txnDC := GetCollection(CollectionFlatTx)

	query := bson.M{
		"hash": txHash,
	}
	var txn *FlatTx
	err := txnDC.Find(query).One(&txn)

	return txn, err
}
