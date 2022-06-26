package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Second)
	//count := 1
	//for _ = range ticker.C {
	//	fmt.Println("ticker..")
	//	count++
	//	if count >= 5 {
	//		ticker.Stop()
	//		break
	//	}
	//
	//}

	chanInt := make(chan int)
	go func() {
		for _ = range ticker.C {
			select {
			case chanInt <- 1:
			case chanInt <- 2:
			case chanInt <- 3:
			}
		}
	}()
	sum := 0
	for c := range chanInt {
		fmt.Printf("收到：%v\n", c)
		sum += c
		if sum >= 15 {
			fmt.Println("完成")
			break
		}
	}
}
