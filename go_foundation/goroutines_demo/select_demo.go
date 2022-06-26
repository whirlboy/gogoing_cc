package main

import (
	"fmt"
	"time"
)

func main() {
	var chanInt = make(chan int)
	var chanStr = make(chan string)

	go func() {
		chanInt <- 100
		chanStr <- "hello"
		defer close(chanInt)
		defer close(chanStr)
	}()
	for {
		select {
		case r := <-chanInt:
			fmt.Printf("chanInt:%v\n", r)
		case r := <-chanStr:
			fmt.Printf("chanStr:%v\n", r)
		default:
			fmt.Println("default...")
		}
		time.Sleep(time.Second)
	}
}
