package db

import (
	"github.com/globalsign/mgo/bson"
	"github.com/iost-official/Go-IOS-Protocol/common"
	"github.com/iost-official/explorer/backend/model/blkchain"
	"log"
)

type Block struct {
	ParentHash  string `bson:"parentHash"`
	Hash        string `bson:"hash"`
	TxsHash     string `bson:"txsHash"`
	MerkleHash  string `bson:"merkleHash"`
	BlockNumber int64  `bson:"blockNumber"`
	TxNumber    int64  `bson:"txNumber"`
	Witness     string `bson:"witness"`
	Time        int64  `bson:"time"`
	Version     int64  `bson:"version"`
	Info        string `bson:"info"`
}

// record failed sync block
type FailBlock struct {
	BlockNumber int64 `bson:"blockNumber"`
	RetryTimes  int64 `bson:"retryTimes"`
	Processed   bool  `bson:"processed"`
}

func GetLastBlockNumber() (int64, error) {
	collection, err := GetCollection("block")
	if err != nil {
		return -1, nil
	}
	var block Block
	err = collection.Find(bson.M{}).Sort("-blockNumber").One(&block)
	if err != nil && err.Error() == "not found" {
		return 0, nil
	}
	return block.BlockNumber, nil
}

func GetTopBlock2() (*Block, error) {
	collection, err := GetCollection(CollectionBlocks)
	if err != nil {
		return nil, err
	}

	var emptyQuery interface{}
	var topBlk *Block
	err = collection.Find(emptyQuery).Sort("-blockNumber").Limit(1).One(&topBlk)
	if err != nil {
		log.Println("getTopBlock error:", err)
		return nil, err
	}

	return topBlk, nil
}

func GetBlockTxnHashes(blockNumber int64) (*[]string, error) {
	txnC, err := GetCollection(CollectionTxs)
	if nil != err {
		log.Println("Get block txn hashes failed", err)
		return nil, err
	}

	var hashes []*struct{ Hash string `bson:"hash"` }
	err = txnC.Find(bson.M{"blockNumber": blockNumber}).Select(bson.M{"hash": 1, "_id": 0}).All(&hashes)
	if nil != err {
		log.Println("query block tx failed", err)
		return nil, err
	}
	result := make([]string, len(hashes))
	for index, hash := range hashes {
		result[index] = hash.Hash
	}
	return &result, nil
}

func GetBlockInfoByNum(num int64) (*Block, *[]string, error) {
	blockInfo, err := blkchain.GetBlockByNum(num, false)

	if nil != err {
		return nil, nil, err
	}

	block := Block{
		ParentHash:  common.Base58Encode(blockInfo.Head.ParentHash),
		Hash:        common.Base58Encode(blockInfo.Hash),
		TxsHash:     common.Base58Encode(blockInfo.Head.TxsHash),
		MerkleHash:  common.Base58Encode(blockInfo.Head.MerkleHash),
		TxNumber:    int64(len(blockInfo.Txhash)),
		BlockNumber: num,
		Witness:     blockInfo.Head.Witness,
		Time:        blockInfo.Head.Time,
		Version:     blockInfo.Head.Version,
		Info:        common.Base58Encode(blockInfo.Head.Info),
	}

	if len(blockInfo.Txhash) > 0 {
		txHashes := make([]string, len(blockInfo.Txhash))
		for i := 0; i < len(blockInfo.Txhash); i++ {
			txHashes[i] = common.Base58Encode(blockInfo.Txhash[i])
		}

		return &block, &txHashes, nil
	}

	return &block, nil, nil
}

func GetBlockByHash(hash string) (*Block, *[]string, error)  {
	blockCollection, err := GetCollection(CollectionBlocks)

	if nil != err {
		log.Println("get block by hash can not get collection", err)
		return nil, nil, err
	}

	var block Block

	err = blockCollection.Find(bson.M{"hash": hash}).One(&block)
	if nil != err {
		log.Println("get block by hash can not find block by hash", err)
		return nil, nil, err
	}

	blockTxHashes, err := GetBlockTxnHashes(block.BlockNumber)

	if nil != err {
		log.Println("get block by hash can not find block tx hashes", err)
		return &block, nil, nil
	}

	return &block, blockTxHashes, nil
}
