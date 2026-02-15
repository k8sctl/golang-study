package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	/*
		A int8 // 1
		B int  // 8
		C int8 // 1
		D int  // 8
		E int8 // 1
		// 총 19바이트이지만 정렬과 패딩 때문에 40바이트의 메모리 공간을 차지한다.

	*/
	A int8 // 1
	C int8 // 1
	E int8 // 1
	B int  // 8
	D int  // 8
	// 총 19바이트이지만 정렬과 패딩 때문에 24바이트의 메모리 공간을 차지한다.
	// A,C,E 뒤에 B를 8바이트 경계에 맞추기 위한 패딩이 들어가고, 마지막 정렬까지 반영되어 최종 24바이트가 된다.
	//즉, 필드 순서만 바꿔도 구조체 크기가 줄어들 수 있다.
}

func main() {
	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))
}

/*
[출력 결과]
24
*/
