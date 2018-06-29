package transport

import (
	"sync"
	"time"

	"gopkg.in/mgo.v2"
)

var (
	mongoSessionMap 		= make(map[string]*mgo.Session)
	mongoSessionMapLock 	sync.RWMutex
)

func GetMongoClient(address string) (*mgo.Session, error) {
	mongoSessionMapLock.RLock()
	if session, ok := mongoSessionMap[address]; ok {
		mongoSessionMapLock.RUnlock()
		return session, nil
	}
	mongoSessionMapLock.RUnlock()

	mongoSessionMapLock.Lock()
	defer mongoSessionMapLock.Unlock()

	session, err := mgo.DialWithTimeout(address, time.Second * 5)
	if err != nil {
		return nil, err
	}

	// to do. close session

	session.SetMode(mgo.Eventual, true)

	mongoSessionMap[address] = session
	return session, nil
}
