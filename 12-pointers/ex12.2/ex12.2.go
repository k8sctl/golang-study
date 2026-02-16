package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	var p1 *int = &a
	var p2 *int = &a
	var p3 *int = &b

	fmt.Printf("p1 == p2: %v\n", p1 == p2)
	fmt.Printf("p2 == p3: %v\n", p2 == p3)
}

/*
[출력 결과]
p1 == p2: true
p2 == p3: false
*/

/*
[코드 동작]
- p1, p2는 모두 변수 a의 주소를 저장하므로 같은 주소를 가리킨다.
- p3는 변수 b의 주소를 저장하므로 p2와 다른 주소를 가리킨다.
- 포인터 비교(==)는 "가리키는 값"이 아니라 "주소값"이 같은지를 비교한다.
- 따라서 p1 == p2 는 true, p2 == p3 는 false가 된다.
*/
