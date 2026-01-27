package main

import "fmt"

func main() {
	// 1. 가장 기본적인 변수 선언 및 할당 형태
	var a int = 3

	// 2. int 타입의 기본 값인 0이 할당된다.
	var b int

	// 3. 정수형 기본 타입 int 타입
	var c = 4

	// 4. var 키워드 생략 및 타입 생략
	// 정수형 기본 타입으로 int 타입으로 자동 지정된다.
	d := 5

	// 타입을 생략했지만 값에 따라 자동으로 String 타입이 지정된다.
	var e = "Hello"

	// 실수형은 float64가 기본 타입이다.
	f := 3.14

	// 선언한 변수들을 출력하여 값이 정상적으로 저장되었는지 확인한다.
	fmt.Println(a, b, c, d, e, f)
}
