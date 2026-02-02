package main

import (
	"fmt"
)

func Add(a int, b int) int {
	// a와 b를 더한 결과를 반환합니다.
	return a + b
}

func main() {
	c := Add(3, 6)
	fmt.Println(c)
}
