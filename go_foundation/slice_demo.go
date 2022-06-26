package main

import "fmt"

func test01() {
	var sliceType1 []int
	sliceType2 := make([]int, 0)
	fmt.Printf("type1:%T,type2:%T\n", sliceType1, sliceType2)
	fmt.Printf("type1:%v,type2:%v\n", sliceType1, sliceType2)
}

func test02() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[0:3]
	fmt.Printf("s1:%v,s2:%v\n", s1, s2)
}

func test03() {
	m1 := map[string]string{"a": "a", "b": "b"}
	for kv := range m1 {
		fmt.Printf("%T\n", kv)
	}
}

func main() {
	test01()
	test02()
	test03()
}
