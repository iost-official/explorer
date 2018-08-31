package dbv2

import (
	"github.com/iost-official/explorer/backend/util/transport"
	"gopkg.in/mgo.v2"
)

const (
	MongoLink = "mongodb://127.0.0.1:27017"
	Db = "explorer"
)

const (
	CollectionBlocks  = "blocks"
	CollectionTxs     = "txs"
	CollectionFlatTx  = "flatxs"
	CollectionFBlocks = "fBlocks"
)

func GetDb() (*mgo.Database, error) {
	mongoClient, err := transport.GetMongoClient(MongoLink)
	if err != nil {
		return nil, err
	}

	return mongoClient.DB(Db), nil
}

func GetCollection(c string) (*mgo.Collection, error) {
	db, err := GetDb()
	if err != nil {
		return nil, err
	}

	return db.C(c), nil
}
