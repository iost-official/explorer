package db

import (
	"github.com/globalsign/mgo/bson"
)

type TaskCursor struct {
	Name   string        `json:"name"`
	Cursor bson.ObjectId `json:"cursor"`
}

const AccountCursorName = "Account_cursor"

func GetAccountTaskCursor() (bson.ObjectId, error) {
	col := GetCollection(CollectionTaskCursor)
	var taskCursor TaskCursor
	err := col.Find(bson.M{"name": AccountCursorName}).One(&taskCursor)
	if err != nil {
		return bson.NewObjectId(), err
	}
	return taskCursor.Cursor, nil
}

func UpdateAccountTaskCursor(cursor bson.ObjectId) error {
	col := GetCollection(CollectionTaskCursor)
	_, err := col.Upsert(bson.M{"name": AccountCursorName}, bson.M{"$set": bson.M{"cursor": cursor}})
	return err
}
