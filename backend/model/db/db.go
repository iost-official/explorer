package db

import (
	"explorer/util/transport"

	"gopkg.in/mgo.v2"
)

const Db = "explorer"

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