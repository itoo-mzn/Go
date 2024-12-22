package main

import (
	"fmt"
)

type Customer struct {
	ID      string
	Balance float64
}

func main() {
	s1 := []int{1, 2, 3}
	// バグる
	// s2 := s1[:2]
	// 完全スライス式で容量も指定することで、s2を変更したときの影響範囲を先頭２つの要素のみに絞ることができる
	s2 := s1[:2:2]
	s3 := append(s2, 10)
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))

	ints := []int{1, 2, 3}
	for i, v := range ints {
		fmt.Println(i, &v, &ints[i])
	}
	fmt.Println(ints)

	bytes := []byte{'a', 'b'}
	strings := []string{"A", "B"}
	fmt.Println(bytes, string(bytes))
	fmt.Println(strings)

	hello := "Hello, 田中, World"
	fmt.Println(hello[:5])
	fmt.Println(hello[:9])

}
