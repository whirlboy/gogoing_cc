package main

import (
	"fmt"
	"runtime"
)

func show(msg string) {
	for i := 0; i < 10; i++ {
		if i >= 5 {
			runtime.Goexit() //runtime.Goexit()：结束当前协程
		}
		fmt.Printf("msg:%v\n", msg)
	}
}

func main() {
	fmt.Printf("runtime.NumCpu():%v\n", runtime.NumCPU())
	go show("java")
	for i := 0; i < 10; i++ {
		runtime.Gosched() //runtime.Gosched()：我有权利执行任务了，腾出CPU让给其他子协程
		fmt.Println("goland..")
	}
	fmt.Println("end...")
}
