package main

import "fmt"

func main() {
	a := 3
	var b float64 = 3.5

	// 1. 변수 b는 float64 타입, int 타입인 변수 c에 저장할 수 없다.
	// var c int = b

	// 타입 변환을 통해 문제를 해결한다.
	var c int = int(b) // 3.5 -> 3

	// 2. 변수 a는 int 타입, 변수 b는 float64 두 변수는 연산이 불가능하다.
	// d := a * b

	// 타입 변환을 통해 문제를 해결한다.
	d := float64(a) * b

	var e int64 = 7

	// 3.
	// 변수 a는 int 타입, 64bit 컴퓨터에서 = int64
	// 변수 e는 명시적으로 int64 타입을 지정하였다.
	// 어떻게 보면 같은 int64 타입을 갖지만 두 변수는 연산이 불가능하다.
	// f := a * e

	// 타입 변환을 통해 문제를 해결한다.
	f := a * int(e)

	fmt.Println(a, b, c, d, e, f)
	// 3 3.5 3 10.5 7 21
}
