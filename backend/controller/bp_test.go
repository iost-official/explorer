package controller

import (
	"fmt"
	"testing"
)

func TestGetChainInfo(t *testing.T) {
	witnessList, err := getChainInfo()
	fmt.Println(witnessList)
	fmt.Println(err)
}

func TestGetBPAccount(t *testing.T) {
	witnessList, err := getBPAccount("4qbWfkLy8ZxmiQXhYYDf6YeJ9N4eNkN8uNEKxWHD4FuG")
	fmt.Println(witnessList)
	fmt.Println(err)
}

func TestGetBPAccountDetail(t *testing.T) {
	ae, err := getBPAccountDetail("vulcanus")
	fmt.Println(ae)
	fmt.Println(err)
}
