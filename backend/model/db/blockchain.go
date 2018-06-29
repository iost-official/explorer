package db

import (
	"log"

	"github.com/iost-official/prototype/rpc"
	"gopkg.in/mgo.v2/bson"
	"encoding/hex"
)

func GetBlocks(start, limit int) ([]*rpc.BlockInfo, error) {
	collection, err := GetCollection("blocks")
	if err != nil {
		log.Println("GetBlock get collection error:", err)
		return nil, err
	}

	var (
		emptyQuery  interface{}
		blkInfoList []*rpc.BlockInfo
	)
	err = collection.Find(emptyQuery).Sort("-head.number").Skip(start).Limit(limit).All(&blkInfoList)
	if err != nil {
		log.Println("GetBlock collection find error:", err)
		return nil, err
	}

	return blkInfoList, nil
}

func GetBlockByLayer(layer int64) (*rpc.BlockInfo, error) {
	// wait implement...
	return nil, nil
}

func GetBlockByHeight(height int64) (*rpc.BlockInfo, error) {
	collection, err := GetCollection("blocks")
	if err != nil {
		return nil, err
	}

	blkQuery := bson.M{
		"head.number": height,
	}
	var blk *rpc.BlockInfo
	err = collection.Find(blkQuery).One(&blk)

	if err != nil {
		return nil, err
	}

	return blk, nil
}

func GetBlockByHash(blkHash string) (*rpc.BlockInfo, error) {
	blkHashDecode, err := hex.DecodeString(blkHash)
	if err != nil {
		return nil, err
	}

	collection, err := GetCollection("blocks")
	if err != nil {
		return nil, err
	}

	blkQuery := bson.M{
		"head.blockhash": blkHashDecode,
	}
	var blk *rpc.BlockInfo
	err = collection.Find(blkQuery).One(&blk)

	if err != nil {
		return nil, err
	}

	return blk, nil
}

func GetTopBlock() (*rpc.BlockInfo, error) {
	collection, err := GetCollection("blocks")
	if err != nil {
		return nil, err
	}

	var emptyQuery interface{}
	var topBlk *rpc.BlockInfo
	err = collection.Find(emptyQuery).Sort("-head.number").Limit(1).One(&topBlk)
	if err != nil {
		log.Println("getTopBlock error:", err)
		return nil, err
	}

	return topBlk, nil
}

func GetBlockLastPage(eachPage int64) int64 {
	var pageLast int64
	if topBlock, err := GetTopBlock(); err == nil {
		if topBlock.Head.Number % eachPage == 0 {
			pageLast = topBlock.Head.Number / eachPage
		} else {
			pageLast = topBlock.Head.Number / eachPage + 1
		}
	}

	return pageLast
}
