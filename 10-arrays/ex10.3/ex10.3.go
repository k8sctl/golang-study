package main

import "fmt"

// 배열 순회

func main() {
	nums := [...]int{10, 20, 30, 40, 50} // [5]int{10, 20, 30, 40, 50}과 같다.

	nums[2] = 300 // 2번 인덱스, 즉 30을 300으로 변경한다.

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}
