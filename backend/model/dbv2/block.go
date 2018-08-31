package dbv2

import (
	"gopkg.in/mgo.v2/bson"
)

type Block struct {
	ParentHash string `bson:"parentHash"`
	TxsHash string `bson:"txsHash"`
	MerkleHash string `bson:"MerkleHash"`
	Number int64 `bson:"number"`
	Witness string `bson:"witness"`
	Time string `bson:"time"`
}

func GetLastBlockNumber () (int64, error){
	collection, err := GetCollection("block")
	if err != nil {
		return -1, nil
	}
	var block Block
	err = collection.Find(bson.M{}).Sort("-number").One(&block)
	if err != nil && err.Error() == "not found" {
		return 0, nil
	}
	return block.Number, nil
}

