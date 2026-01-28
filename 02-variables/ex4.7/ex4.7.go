package main

import "fmt"

func main() {
	// 실수 계산시 오차
	var a float32 = 1234.523
	var b float32 = 3456.123
	var c float32 = a * b
	var d float32 = c * 3

	fmt.Println(a) // 1234.523
	fmt.Println(b) // 3456.123
	fmt.Println(c) // 4.266663e+06, 실제 원래 결과값은 4266663.334329
	fmt.Println(d) // 1.2799989e+07, 실제 원래 결과값은 12799990.002987

	// float32는 7자리
	// float64는 15자리
}
