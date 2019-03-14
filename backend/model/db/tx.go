package db

import (
	"log"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/iost-official/explorer/backend/model/blockchain/rpcpb"
)

type ActionRaw struct {
	Contract   string `bson:"contract" json:"contract"`
	ActionName string `bson:"actionName" json:"actionName"`
	Data       string `bson:"data" json:"data"`
}

type SignatureRaw struct {
	Algorithm int32  `bson:"algorithm" json:"algorithm"`
	Sig       string `bson:"sig" json:"sig"`
	PubKey    string `bson:"pubKey" json:"pubKey"`
}

type ReceiptRaw struct {
	Type    int32  `bson:"type" json:"type"`
	Content string `bson:"content" json:"content"`
}

type TxReceiptRaw struct {
	GasUsage      int64        `bson:"gasUsage"`
	SuccActionNum int32        `bson:"succActionNum"`
	Receipts      []ReceiptRaw `bson:"receipts"`
	StatusCode    int32        `bson:"statusCode"`
	StatusMessage string       `bson:"statusMessage"`
}

type TmpTx struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Hash        string        `bson:"hash"`
	BlockNumber int64         `bson:"blockNumber"`
	Mark        int64         `bson:"mark"`
}

type Tx struct {
	ExternalId  bson.ObjectId  `bson:"externalId"`
	BlockNumber int64          `bson:"blockNumber"`
	Time        int64          `bson:"time"`
	Hash        string         `bson:"hash"`
	Expiration  int64          `bson:"expiration"`
	GasPrice    int64          `bson:"gasPrice"`
	GasLimit    int64          `bson:"gasLimit"`
	Mark        int64          `bson:"mark"`
	Actions     []ActionRaw    `bson:"actions"`
	Signers     []string       `bson:"signers"`
	Signs       []SignatureRaw `bson:"signs"`
	Publisher   SignatureRaw   `bson:"publisher"`
	Receipt     TxReceiptRaw   `bson:"receipt"`
}

type TxStore struct {
	BlockNumber int64              `json:"block_number"`
	Tx          *rpcpb.Transaction `json:"tx"`
}

// 将 Tx.Actions 打平后的数据结构， 如果actionName == Transfer 则会解析出 from, to, amount
type FlatTx struct {
	Id          bson.ObjectId  `bson:"_id,omitempty" json:"id"`
	BlockNumber int64          `bson:"blockNumber" json:"blockNumber"`
	Time        int64          `bson:"time" json:"time"`
	Hash        string         `bson:"hash" json:"hash"`
	Expiration  int64          `bson:"expiration" json:"expiration"`
	GasPrice    int64          `bson:"gasPrice" json:"gasPrice"`
	GasLimit    int64          `bson:"gasLimit" json:"gasLimit"`
	Action      ActionRaw      `bson:"action" json:"action"`
	Signers     []string       `bson:"signers" json:"signers"`
	Signs       []SignatureRaw `bson:"signs" json:"signs"`
	Publisher   string         `bson:"publisher" json:"publisher"`
	From        string         `bson:"from" json:"from"`
	To          string         `bson:"to" json:"to"`
	Amount      float64        `bson:"amount" json:"amount"`           // 转发数量
	ActionIndex int            `bson:"actionIndex" json:"actionIndex"` // action 索引
	ActionName  string         `bson:"actionName" json:"actionName"`   // action 类型
	Receipt     TxReceiptRaw   `bson:"receipt" json:"receipt"`
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		log.Printf("%s took %v\n", what, time.Since(start))
	}
}

func ProcessTxs(txs []*rpcpb.Transaction, blockNumber int64) error {
	insertTxs(txs, blockNumber)
	return nil
}

func insertTxs(txs []*rpcpb.Transaction, blockNumber int64) {
	var txnC *mgo.Collection
	txnC = GetCollection(CollectionTxs)

	for _, tx := range txs {
		txStore := TxStore{BlockNumber: blockNumber, Tx: tx}
		for {
			_, err := txnC.Upsert(bson.M{"tx.hash": tx.Hash}, txStore)
			if err != nil {
				log.Println("fail to insert txs, err: ", err)
				time.Sleep(time.Second)
				continue
			} else {
				//log.Println("update txs, txHash: ", tx.Hash)
				break
			}
		}
	}
	log.Println("update txs, size: ", len(txs))
	/*txInterfaces := make([]interface{}, len(txs))
	for i, tx := range txs {
		txInterfaces[i] = TxStore{BlockNumber: blockNumber, Tx: tx}
	}

	for {
		err := txnC.Insert(txInterfaces...)
		if err != nil && strings.Index(err.Error(), "duplicate key") == -1 {
			log.Println("fail to insert txs, err: ", err)
			time.Sleep(time.Second)
			continue
		} else {
			log.Println("update txs, size: ", len(txs))
			break
		}
	}*/
}

func GetTxByHash(hash string) (*TxStore, error) {
	txnDC := GetCollection(CollectionTxs)
	query := bson.M{
		"tx.hash": hash,
	}
	var tx *TxStore
	err := txnDC.Find(query).One(&tx)

	return tx, err
}

func GetTxsByHash(hashes []string) ([]*TxStore, error) {
	txnDC := GetCollection(CollectionTxs)
	query := bson.M{
		"tx.hash": bson.M{
			"$in": hashes,
		},
	}
	var txs []*TxStore
	err := txnDC.Find(query).All(&txs)
	if err != nil {
		return nil, err
	}

	txMap := make(map[string]*TxStore)
	for _, t := range txs {
		txMap[t.Tx.Hash] = t
	}

	ret := make([]*TxStore, 0, len(txMap))
	for _, hash := range hashes {
		ret = append(ret, txMap[hash])
	}

	return ret, err
}

func GetTxCountByNumber(number int64) (int, error) {
	txnDC := GetCollection(CollectionTxs)
	query := bson.M{}
	if number >= 0 {
		query["blocknumber"] = number
	}
	return txnDC.Find(query).Count()
}

// ConvertTxs used to convert tx in db to web display format
func convertTxs(txs []*rpcpb.Transaction) []FlatTx {

	return nil
}

func GetTxTotalPageCnt(eachPage, blockNumber int64) (int64, error) {
	txnDC := GetCollection(CollectionTxs)
	query := bson.M{}
	if blockNumber >= 0 {
		query["blocknumber"] = blockNumber
	}

	count, err := txnDC.Find(query).Count()
	if err != nil {
		return 0, err
	}
	txsInt64Len := int64(count)
	pageMax := txsInt64Len / eachPage
	if txsInt64Len%eachPage != 0 {
		pageMax++
	}
	return pageMax, nil
}

func GetTxs(blockNumber int64, start, limit int) ([]*TxStore, error) {
	txnDC := GetCollection(CollectionTxs)
	query := bson.M{}
	if blockNumber >= 0 {
		query["blocknumber"] = blockNumber
	}
	var txStores []*TxStore

	err := txnDC.Find(query).Sort("-tx.time").Skip(start).Limit(limit).All(&txStores)
	if err != nil {
		return nil, err
	}
	return txStores, nil
}
