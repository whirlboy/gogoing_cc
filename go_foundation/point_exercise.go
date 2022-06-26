package main

import "fmt"

func point_test01() {
	a := [3]int{1, 2, 3}
	var pa [3]*int
	fmt.Printf("a:%T,%v\n", a, a)
	fmt.Printf("pa:%T,%v\n", pa, pa)
}

func main() {
	point_test01()
}
