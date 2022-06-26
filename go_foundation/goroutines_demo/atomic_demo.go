package main

import "sync/atomic"



func main() {
	var i = 100

	func add(){
		atomic.AddInt32(&i,1)
	}

	func sub(){

	}
}
