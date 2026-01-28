package main

import "fmt"

func main() {
	var a int16 = 3456   // a는 int16 타입, 2바이트 정수
	var b int8 = int8(a) // int16 -> int8

	fmt.Println(a, b) // 3456 -128

	// 왜 b의 값이 a의 값과 다를까?
	// int16은 2바이트, int8은 1바이트 사이즈를 갖기 때문에
	// 타입 변환을 할 때, 1바이트를 넘어가는 부분은 버려진다.

	// 예시:
	// a int16 = 3456 = 00001101 10000000(2진수)
	// b int8 = int(a) = -------- 10000000(2진수)
	// 10000000을 int8로 표현하면 -128이 된다.
}
