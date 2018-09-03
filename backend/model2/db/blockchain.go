package db

import (
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

	err = blockCollection.Find(emptyQuery).Sort("-head.number").Skip(start).Limit(limit).All(&blkInfoList)

	if nil != err {
		log.Println("Get blocks collection query err", err)
		return nil, err
	}

	return blkInfoList, nil
}
