package main

import (
	"fmt"
	"unsafe"
)

// 필드 배치에 따른 구조체 크기 변화

type User struct {
	Age   int32   // 4바이트
	Score float64 // 8바이트
}

func main() {
	user := User{32, 77.2}
	fmt.Println(unsafe.Sizeof(user))
	// 구조체 User가 메모리에서 차지하는 바이트 크기를 반환
	// int32(4바이트) + float64(8바이트)라 단순 합은 12바이트지만,
	// 실제로는 정렬(alignment)과 패딩(padding) 때문에 16바이트가 된다.
}

/*
[출력 결과]
16
*/
