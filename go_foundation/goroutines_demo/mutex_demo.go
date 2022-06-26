package main

import (
	"fmt"
	"sync"
	"time"
)

var i int = 100

var wt sync.WaitGroup

var lock sync.Mutex

func add() {
	lock.Lock()
	defer wt.Done()
	i += 1
	fmt.Printf("add,%v\n", i)
	time.Sleep(time.Millisecond * 10)
	lock.Unlock()
}

func sub() {
	lock.Lock()
	time.Sleep(time.Millisecond * 2)
	defer wt.Done()
	i -= 1
	fmt.Printf("sub,%v\n", i)
	lock.Unlock()

}

func main() {
	for i := 0; i < 100; i++ {
		wt.Add(1)
		go add()
		wt.Add(1)
		go sub()
	}
	wt.Wait()
	fmt.Printf("finnal:%v\n", i)
}
