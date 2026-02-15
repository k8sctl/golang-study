package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8 // 1
	B int  // 8
	C int8 // 1
	D int  // 8
	E int8 // 1
	// 총 19바이트
}

func main() {
	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))
}

/*
[출력 결과]
40
*/
