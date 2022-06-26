package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 2)
	fmt.Printf("time.Now():%v\n", time.Now())
	t1 := <-timer.C //阻塞的时间
	fmt.Printf("t1:%v\n", t1)
}
