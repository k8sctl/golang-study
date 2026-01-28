package main

import "fmt"

var g int = 10

func main() {
	var m int = 20

	{
		var s int = 50
		fmt.Println(m, s, g)
	}

	// m = s + 20
	// 변수의 범위는 그 변수가 속한 {}(중괄호, 블록) 안 까지이다.
	// 즉 변수 s를 선언한 코드 블록 외부에서 변수 s를 사용하려고 하기 때문에
	// undefined: s 에러 메시지를 출력한다.
	// 또한 변수 g는 패키지 전체에 살아있는 변수이기 때문에 '패키지 전역변수'라고 부른다.
}
