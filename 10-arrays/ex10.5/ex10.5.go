package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 200, 100}

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	fmt.Println()

	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}

	b = a
	// 같은 타입/길이 배열이라 통째로 복사됨(요소 하니씩 복사되는 것이 아니다.)
	// 사이즈가 다르면 assignment 할 수 없다는 에러가 발생한다.
	// 타입이 달라도 assignment 할 수 없다는 에러가 발생한다.

	fmt.Println()

	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}
}

/*
[출력 결과]
a[0] = 1
a[1] = 2
a[2] = 3
a[3] = 4
a[4] = 5

b[0] = 500
b[1] = 400
b[2] = 300
b[3] = 200
b[4] = 100

b[0] = 1
b[1] = 2
b[2] = 3
b[3] = 4
b[4] = 5
*/
