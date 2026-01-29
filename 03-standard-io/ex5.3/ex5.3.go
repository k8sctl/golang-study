package main

import "fmt"

func main() {

	/*
		[표준 입력]
		Scan() : 표준 입력에서 값을 입력받는다.
		Scanf() : 표준 입력에서 서식 형태로 값을 입력받는다.
		Scanln() : 표준 입력에서 한 줄을 읽어서 값을 입력받는다.
	*/

	var a int // 정수 타입 기본값 0
	var b int // 정수 타입 기본값 0

	n, err := fmt.Scanln(&a, &b)
	// 변수에 '&'기호를 붙히면 해당 변수의 메모리 주소값을 의미한다.

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}

	/*
		입력 : 10 20
		출력 : 2 10 20
	*/
}
