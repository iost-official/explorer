package db

import (
	"github.com/globalsign/mgo/bson"
	"log"
)

func GetBlocks(start, limit int) ([]*Block, error) {
	blockCollection, err := GetCollection(CollectionBlocks)
	if nil != err {
		log.Println("Get blocks get blocks db err", err)
		return nil, err
	}
	var (
		emptyQuery  interface{}
		blkInfoList []*Block
	)

	err = blockCollection.Find(emptyQuery).Sort("-blockNumber").Skip(start).Limit(limit).All(&blkInfoList)

	if nil != err {
		log.Println("Get blocks collection query err", err)
		return nil, err
	}

	return blkInfoList, nil
}

func GetTopBlock() (*Block, error) {
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

func GetBlockLastPage(eachPage int64) int64 {
	var pageLast int64
	if topBlock, err := GetTopBlock(); err == nil {
		if topBlock.BlockNumber%eachPage == 0 {
			pageLast = topBlock.BlockNumber / eachPage
		} else {
			pageLast = topBlock.BlockNumber/eachPage + 1
		}
	}

	return pageLast
}

func GetBlockByHeight(height int64) (*Block, error) {
	collection, err := GetCollection(CollectionBlocks)
	if err != nil {
		return nil, err
	}

	blkQuery := bson.M{
		"blockNumber": height,
	}
	var blk *Block
	err = collection.Find(blkQuery).One(&blk)

	if err != nil {
		return nil, err
	}

	return blk, nil
}
