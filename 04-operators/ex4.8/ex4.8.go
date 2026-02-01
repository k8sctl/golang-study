package main

import (
	"fmt"
	"math"
)

func equal(a, b float64) bool {
	// 같으면 true
	// 다르면 false를 반환
	return math.Nextafter(a, b) == b
}

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, equal(a+b, c))
	/*
		[출력]
		0.299999999999999989 == 0.300000000000000044 : true
	*/

	// math.Nextafter(A, B)는
	// A -> B 방향으로 1비트만큼 조정한다.
}
