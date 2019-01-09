package db

import "time"

type ApplyAccount struct {
	PubKey     string
	Name       string
	Email      string
	RemoteIP   string
	InsertTime int64
}

func AddApply(accountPubKey, accountName, email, remoteIP string) error {
	applyIOSTC := GetCollection(CollectionApplyIOST)

	record := &ApplyAccount{
		PubKey:     accountPubKey,
		Name:       accountName,
		Email:      email,
		RemoteIP:   remoteIP,
		InsertTime: time.Now().Unix(),
	}

	return applyIOSTC.Insert(record)
}
