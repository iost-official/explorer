package db

import (
	"log"
	"time"

	"github.com/globalsign/mgo"
	"github.com/iost-official/iost-api/util/transport"
)

func GetDb() (*mgo.Database, error) {
	var err error
	var mongoClient *mgo.Session

	if MongoUser == "" && MongoPassWord == "" {
		mongoClient, err = transport.GetMongoClient(MongoLink, Db)
	} else {
		mongoClient, err = transport.GetMongoClientWithAuth(MongoLink, MongoUser, MongoPassWord, Db)
	}
	if err != nil {
		return nil, err
	}

	return mongoClient.DB(Db), nil
}

func GetCollection(c string) *mgo.Collection {
	var d *mgo.Database
	var err error
	var retryTime int
	for {
		d, err = GetDb()
		if err != nil {
			log.Println("fail to get db collection ", err)
			time.Sleep(time.Second)
			retryTime++
			if retryTime > 10 {
				log.Fatalln("fail to get db collection, retry time exceeds")
			}
			continue
		}
		return d.C(c)
	}
}
