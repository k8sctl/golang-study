package main

import "fmt"

/*
[Go switch의 동작]
- 다른 언어(C, Java 등)는 case 끝에 break를 안 쓰면 다음 case로 계속 진행(fall-through)될 수 있다.
- Go는 기본적으로 해당 case만 실행한 뒤 switch를 자동 종료한다. (즉, break를 보통 직접 쓰지 않아도 됨)

[다음 case도 이어서 실행하고 싶을 때]
- fallthrough 키워드를 사용하면 바로 다음 case를 추가로 실행할 수 있다.
- fallthrough는 "다음 case의 조건 검사 없이" 실행을 넘긴다.
*/

func main() {
	a := 3

	switch a {
	case 1:
		fmt.Println("a == 1")
		break
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
		fallthrough
	case 4:
		fmt.Println("a == 4")
	default:
		fmt.Println(" a != 1, 2, 3, 4")
	}
}
