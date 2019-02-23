package controller

import (
	"fmt"
	"testing"
	"time"

	"github.com/iost-official/explorer/backend/config"
	"github.com/iost-official/explorer/backend/model/db"
)

func TestGetBPAccount(t *testing.T) {
	witnessList, err := getBPAccount("4qbWfkLy8ZxmiQXhYYDf6YeJ9N4eNkN8uNEKxWHD4FuG")
	fmt.Println(witnessList)
	fmt.Println(err)
}

func TestGetBPProduce(t *testing.T) {
	config.ReadConfig("")
	witnessList, err := db.GetBPProduce([]string{"6sNQa7PV2SFzqCBtQUcQYJGGoU7XaB6R4xuCQVXNZe6b"}, time.Now(), time.Now())
	fmt.Println(witnessList)
	fmt.Println(err)
}
func TestGetBPAccountDetail(t *testing.T) {
	ae, err := getBPAccountDetail("vulcanus")
	fmt.Println(ae)
	fmt.Println(err)
}
