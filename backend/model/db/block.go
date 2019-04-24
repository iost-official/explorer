package db

import (
	"log"

	"github.com/globalsign/mgo/bson"
	"github.com/iost-official/explorer/backend/model/blockchain/rpcpb"
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

func GetFirstBlockNumberAfter(time int64) (int64, error) {
	collection := GetCollection(CollectionBlocks)
	var block *rpcpb.Block
	err := collection.Find(bson.M{"time": bson.M{"$gt": time}}).Sort("time").One(&block)
	if err != nil && err.Error() == "not found" {
		return 0, nil
	}
	return block.Number, nil
}

func GetLastBlockNumberBefore(time int64) (int64, error) {
	collection := GetCollection(CollectionBlocks)
	var block *rpcpb.Block
	err := collection.Find(bson.M{"time": bson.M{"$lt": time}}).Sort("-time").One(&block)
	if err != nil && err.Error() == "not found" {
		return 0, nil
	}
	return block.Number, nil
}

func GetLastBlockNumber() (int64, error) {
	collection := GetCollection("block")
	var block *rpcpb.Block
	err := collection.Find(bson.M{}).Sort("-number").Limit(1).One(&block)
	if err != nil && err.Error() == "not found" {
		return 0, nil
	}
	return block.Number, nil
}

func GetBlockTxnHashes(blockNumber int64) (*[]string, error) {
	txnC := GetCollection(CollectionTxs)

	var hashes []*struct {
		Hash string `bson:"hash"`
	}
	err := txnC.Find(bson.M{"blockNumber": blockNumber}).Select(bson.M{"hash": 1, "_id": 0}).All(&hashes)
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

/*func GetBlockInfoByNum(num int64) (*Block, *[]string, error) {
	blockInfo, err := blockchain.GetBlockByNum(num, false)

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
}*/

func GetBlockByHash(hash string) (*rpcpb.Block, *[]string, error) {
	blockCollection := GetCollection(CollectionBlocks)

	var block rpcpb.Block

	err := blockCollection.Find(bson.M{"hash": hash}).One(&block)
	if nil != err {
		log.Println("get block by hash can not find block by hash", err)
		return nil, nil, err
	}
	return &block, nil, nil

	/*  blockTxHashes, err := GetBlockTxnHashes(block.Number) */

	// if nil != err {
	// log.Println("get block by hash can not find block tx hashes", err)
	// return &block, nil, nil
	// }

	/* return &block, blockTxHashes, nil */
}

func GetBlocks(start, limit int) ([]*rpcpb.Block, error) {
	blockCollection := GetCollection(CollectionBlocks)
	var (
		emptyQuery  interface{}
		blkInfoList []*rpcpb.Block
	)

	err := blockCollection.Find(emptyQuery).Sort("-number").Skip(start).Limit(limit).All(&blkInfoList)

	if nil != err {
		log.Println("Get blocks collection query err", err)
		return nil, err
	}

	return blkInfoList, nil
}

func GetTopBlock() (*rpcpb.Block, error) {
	collection := GetCollection(CollectionBlocks)

	var emptyQuery interface{}
	var topBlk *rpcpb.Block
	err := collection.Find(emptyQuery).Sort("-number").Limit(1).One(&topBlk)
	if err != nil {
		log.Println("getTopBlock error:", err)
		return nil, err
	}

	return topBlk, nil
}

func GetBlockLastPage(eachPage int64) int64 {
	var pageLast int64
	if topBlock, err := GetTopBlock(); err == nil {
		if topBlock.Number%eachPage == 0 {
			pageLast = topBlock.Number / eachPage
		} else {
			pageLast = topBlock.Number/eachPage + 1
		}
	}

	return pageLast
}

func GetBlockByHeight(height int64) (*rpcpb.Block, error) {
	collection := GetCollection(CollectionBlocks)

	blkQuery := bson.M{
		"number": height,
	}
	var blk *rpcpb.Block
	err := collection.Find(blkQuery).One(&blk)

	if err != nil {
		return nil, err
	}

	return blk, nil
}
