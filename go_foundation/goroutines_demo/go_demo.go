package main

import (
	"fmt"
	"time"
)

func showMsg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go showMsg("yyy")
	go showMsg("zzz")
	showMsg("ccc")
}
