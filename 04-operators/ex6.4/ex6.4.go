package main

import (
	"fmt"
)

func main() {
	var x int8 = 16
	var y int8 = -128
	var z int8 = -1
	var w uint8 = 128

	fmt.Printf("x:%08b x>>2:%08b x>>2:%d\n", x, x>>2, x>>2)
	fmt.Printf("y:%08b y>>2:%08b y>>2:%d\n", uint8(y), uint8(y>>2), y>>2)
	fmt.Printf("z:%08b z>>2:%08b z>>2:%d\n", uint8(z), uint8(z>>2), z>>2)
	fmt.Printf("w:%08b w>>2:%08b w>>2:%d\n", w, w>>2, w>>2)

	/*
		[츨력]
		x:00010000 x>>2:00000100 x>>2:4 // 양수이기 때문에 시프트 연산을 하면 0이 채워진다.
		y:10000000 y>>2:11100000 y>>2:-32 // 음수이기 때문에 시프트 연산을 하면 1이 채워진다.
		z:11111111 z>>2:11111111 z>>2:-1 // 음수이기 때문에 시프트 연산을 하면 1이 채워진다.
		w:10000000 w>>2:00100000 w>>2:32 // uint8, 부호가 없기 때문에 시프트 연산을 해도 0이 채워진다.
	*/
}
