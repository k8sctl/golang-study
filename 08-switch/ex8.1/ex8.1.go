package main

import "fmt"

/*
	[switch 구문 정리]

	switch 비굣값 { 	// 검사하는 값이 온다.
	case 값1: 		   // 비굣값과 값1이 같을 때 수행한다.
  		문장
	case 값2: 		   // 비굣값과 값2가 같을 때 수행한다.
  		문장
	default:          // 만족하는 case가 없을 때 수행한다.(선택사항)
  		문장
	}
*/

func main() {
	a := 3

	switch a {
	case 1:
		fmt.Println("a == 1")
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
	default:
		fmt.Println("a != 1, 2, 3", a)
	}
}
