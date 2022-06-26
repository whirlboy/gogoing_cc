package main

import "fmt"

//闭包1
func add() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

//闭包VIP
func cal(base int) (func(int) int, func(int) int) {
	add := func(a int) int {
		base += a
		return base
	}
	sub := func(a int) int {
		base -= a
		return base
	}
	return add, sub
}

func main() {
	add, sub := cal(10)
	fmt.Println(add(1), sub(2))
}
