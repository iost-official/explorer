package transport

import (
	"fmt"
	"sync"
	"time"

	"github.com/globalsign/mgo"
)

var (
	mongoSessionMap     = make(map[string]*mgo.Session)
	mongoSessionMapLock sync.RWMutex
)

func GetMongoClientWithAuth(address, username, password, db string) (*mgo.Session, error) {
	mongoSessionMapLock.RLock()
	if session, ok := mongoSessionMap[address]; ok {
		mongoSessionMapLock.RUnlock()
		return session, nil
	}
	mongoSessionMapLock.RUnlock()

	mongoSessionMapLock.Lock()
	defer mongoSessionMapLock.Unlock()

	dInfo := mgo.DialInfo{
		Addrs:    []string{address},
		Timeout:  time.Second * 5,
		Username: username,
		Password: password,
		Database: db,
	}
	session, err := mgo.DialWithInfo(&dInfo)
	if err != nil {
		return nil, err
	}
	fmt.Println("Dial Correct!")

	// to do. close session

	session.SetMode(mgo.Eventual, true)
	session.SetSocketTimeout(time.Minute)

	mongoSessionMap[address] = session
	return session, nil
}

func GetMongoClient(address, db string) (*mgo.Session, error) {
	mongoSessionMapLock.RLock()
	if session, ok := mongoSessionMap[address]; ok {
		mongoSessionMapLock.RUnlock()
		return session, nil
	}
	mongoSessionMapLock.RUnlock()

	mongoSessionMapLock.Lock()
	defer mongoSessionMapLock.Unlock()

	dInfo := mgo.DialInfo{
		Addrs:    []string{address},
		Timeout:  time.Second * 5,
		Database: db,
	}
	session, err := mgo.DialWithInfo(&dInfo)
	if err != nil {
		return nil, err
	}
	fmt.Println("Dial Correct!")

	// to do. close session

	session.SetMode(mgo.Eventual, true)
	session.SetSocketTimeout(time.Minute)

	mongoSessionMap[address] = session
	return session, nil
}
