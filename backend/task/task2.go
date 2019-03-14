package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	time.Sleep(time.Second * 2)

	s1 := time.Since(now)

	time.Sleep(time.Second * 2)

	s2 := time.Since(now)

	fmt.Println((s1 + s2) / 10)

	var a time.Duration
	fmt.Println(a)
}
