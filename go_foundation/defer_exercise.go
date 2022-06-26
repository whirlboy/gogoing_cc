package main

import "fmt"

func defer_test01() {
	fmt.Println("11111")
	defer fmt.Println("22222")
	defer fmt.Println("33333")
	fmt.Println("44444")
}

func main() {
	defer_test01()
}
