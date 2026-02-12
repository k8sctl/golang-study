package main

import (
	"fmt"
)

func main() {
	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9},
	}

	for _, arr := range a {
		for _, v := range arr {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}

/*
[2차원 배열 순회 예제]
- a는 [2][5]int 타입의 2차원 배열이다.
  -> 행(row) 2개, 열(column) 5개를 의미한다.
- 바깥 for-range는 각 행([5]int)을 순회한다.
- 안쪽 for-range는 해당 행의 각 원소(int)를 순회한다.
- fmt.Print(v, " ")로 한 줄에 값들을 출력하고,
  각 행이 끝날 때 fmt.Println()으로 줄바꿈한다.

[출력 결과]
1 2 3 4 5
5 6 7 8 9
*/
