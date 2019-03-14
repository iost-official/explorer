package db

import "time"

type ApplyAccount struct {
	TxHash     string
	PubKey     string
	Name       string
	Email      string
	RemoteIP   string
	InsertTime int64
}

func AddApply(txHash, accountPubKey, accountName, email, remoteIP string) error {
	applyIOSTC := GetCollection(CollectionApplyIOST)

	record := &ApplyAccount{
		TxHash:     txHash,
		PubKey:     accountPubKey,
		Name:       accountName,
		Email:      email,
		RemoteIP:   remoteIP,
		InsertTime: time.Now().Unix(),
	}

	return applyIOSTC.Insert(record)
}
