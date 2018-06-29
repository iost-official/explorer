package db

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	"github.com/iost-official/prototype/common"
	"github.com/iost-official/prototype/core/tx"
	"github.com/iost-official/prototype/rpc"
	"github.com/iost-official/prototype/vm/lua"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"strconv"
	"strings"
)

type ExplorerTx struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	BlockHeight int64         `json:"block_height"`
	Tx          *tx.Tx        `json:"tx"`
	TxHash      []byte        `json:"tx_hash"`
}

type MgoMiddleTx struct {
	Time      int64              `json:"time"`
	Nonce     int64              `json:"nonce"`
	Contract  []byte             `json:"contract"`
	Signs     []common.Signature `json:"signs"`
	Publisher common.Signature   `json:"publisher"`
	Recorder  common.Signature   `json:"recorder"`
}

// to do...
func (et *ExplorerTx) GetBSON() (interface{}, error) {
	encoded := new(struct {
		BlockHeight int64        `json:"block_height"`
		Tx          *MgoMiddleTx `json:"tx"`
		TxHash      []byte       `json:"tx_hash"`
	})

	encoded.BlockHeight = et.BlockHeight
	encoded.TxHash = et.TxHash
	if et.Tx == nil {
		return encoded, nil
	}

	if luaContract, ok := et.Tx.Contract.(*lua.Contract); ok {
		luaContractEncoded := luaContract.Encode()
		encoded.Tx = &MgoMiddleTx{
			Time:      et.Tx.Time,
			Nonce:     et.Tx.Nonce,
			Contract:  luaContractEncoded,
			Signs:     et.Tx.Signs,
			Publisher: et.Tx.Publisher,
			Recorder:  et.Tx.Recorder,
		}
		return encoded, nil
	} else {
		return nil, errors.New("error convert")
	}
}

// to do...
func (et *ExplorerTx) SetBSON(raw bson.Raw) error {
	decoded := new(struct {
		Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
		BlockHeight int64         `json:"block_height"`
		Tx          *MgoMiddleTx  `json:"tx"`
		TxHash      []byte        `json:"tx_hash"`
	})

	bsonErr := raw.Unmarshal(decoded)
	if bsonErr == nil {
		et.Id = decoded.Id
		et.BlockHeight = decoded.BlockHeight
		et.TxHash = decoded.TxHash
		if decoded.Tx == nil {
			return nil
		}

		if len(decoded.Tx.Contract) > 0 {
			luaContract := new(lua.Contract)
			err := luaContract.Decode(decoded.Tx.Contract)
			if err != nil {
				return err
			}

			et.Tx = &tx.Tx{
				Time:      decoded.Tx.Time,
				Nonce:     decoded.Tx.Nonce,
				Contract:  luaContract,
				Signs:     decoded.Tx.Signs,
				Publisher: decoded.Tx.Publisher,
				Recorder:  decoded.Tx.Recorder,
			}

			return nil
		} else {
			return errors.New("err convert")
		}
	} else {
		return bsonErr
	}
}

func (et *ExplorerTx) GenerateMgoTx() *MgoTx {
	if et == nil || et.Tx == nil {
		return nil
	}

	contractInfo := et.Tx.Contract.Info()

	contractCode := et.Tx.Contract.Code()
	codeList := strings.Split(contractCode, `"`)
	var (
		to string
		amount   float64
	)
	if len(codeList) >= 5 {
		//if len(codeList[1]) != 0 {
			//from = codeList[1]
		//}
		if len(codeList[3]) != 0 {
			to = codeList[3]
		}
		if len(codeList[4]) > 5 {
			if bucketIndex := strings.Index(codeList[4], ")"); bucketIndex >= 0 {
				amountStr := codeList[4][1:bucketIndex]
				amount, _ = strconv.ParseFloat(amountStr, 64)
			}
		}
	}

	return &MgoTx{
		MgoSourceId: et.Id,
		BlockHeight: et.BlockHeight,
		TxHash:      et.Tx.Hash(),
		Time:        et.Tx.Time,
		Nonce:       et.Tx.Nonce,
		Publisher:   et.Tx.Publisher,
		GasLimit:    contractInfo.GasLimit,
		Price:       contractInfo.Price,
		From:        common.Base58Encode(et.Tx.Publisher.Pubkey),
		To:          to,
		Amount:      amount,
		Code:        contractCode,
	}
}

func GetTxn(start, limit int) ([]*ExplorerTx, error) {
	txnC, err := GetCollection("txns")
	if err != nil {
		log.Println("UpdateTxns get txns collection error:", err)
		return nil, err
	}

	txnQuery := bson.M{
		"tx": bson.M{"$ne": nil},
	}
	var txns []*ExplorerTx
	err = txnC.Find(txnQuery).Sort("-_id").Skip(start).Limit(limit).All(&txns)
	if err != nil {
		return nil, err
	}

	return txns, nil
}

func GetTxnByKey(tkey *rpc.TransactionKey) (*ExplorerTx, error) {
	txnC, err := GetCollection("txns")
	if err != nil {
		log.Println("UpdateTxns get txns collection error:", err)
		return nil, err
	}

	txnQuery := bson.M{
		"tx.nonce":            tkey.Nonce,
		"tx.publisher.pubkey": tkey.Publisher,
	}

	var txn *ExplorerTx
	err = txnC.Find(txnQuery).One(&txn)
	if err != nil {
		return nil, err
	}

	return txn, err
}

func GetTopTxn() (*ExplorerTx, error) {
	txnC, err := GetCollection("txns")
	if err != nil {
		log.Println("UpdateTxns get txns collection error:", err)
		return nil, err
	}

	var (
		emptyQuery interface{}
		topTxn     *ExplorerTx
	)
	err = txnC.Find(emptyQuery).Sort("-_id").Limit(1).One(&topTxn)
	if err != nil {
		log.Println("getTopTxn error:", err)
		return nil, err
	}

	return topTxn, nil
}

func GetTotalTxnLen() (int, error) {
	txnC, err := GetCollection("txns")
	if err != nil {
		log.Println("UpdateTxns get txns collection error:", err)
		return 0, err
	}

	totalQuery := bson.M{
		"tx": bson.M{"$ne": nil},
	}
	return txnC.Find(totalQuery).Count()
}

func GetTxLastPage(eachPage int64) (int64, error) {
	totalLen, err := GetTotalTxnLen()
	if err != nil {
		return 0, err
	}
	txsInt64Len := int64(totalLen)

	var pageLast int64
	if txsInt64Len%eachPage == 0 {
		pageLast = txsInt64Len / eachPage
	} else {
		pageLast = txsInt64Len / eachPage
	}

	return pageLast, nil
}
