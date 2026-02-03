package main

import (
	"fmt"
)

func main() {
	const pi1 float64 = 3.141592653589793228 // 상수
	var pi2 float64 = 3.141592653589793228   // 변수

	// pi1 = 3
	// // pi1은 상수로 선언했기 때문에 값을 변경하려고 하면
	// 'cannot assign to pi1(declared const)' 에러가 발생한다.
	pi2 = 4

	fmt.Printf("원주율: %f\n", pi1)
	fmt.Printf("원주율: %f\n", pi2)

	/*
		[출력]
		원주율: 3.141593
		원주율: 4.000000
	*/
}
