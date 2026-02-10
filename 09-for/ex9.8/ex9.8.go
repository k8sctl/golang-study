package main

import "fmt"

/*
[중첩 for문 종료 방법]
중첩 for문에서 특정 조건이 만족되면 바깥 반복문까지 종료해야 할 때는 다음 방법을 사용한다.

2) 레이블(label) 사용
   - 바깥 루프에 레이블을 붙이고 `break 레이블명`으로
   - 원하는 반복문을 즉시 종료한다.
*/

func main() {
	a := 1
	b := 1

OuterFor: // Label, 레이블
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				break OuterFor
			}
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}

/*
	[출력 결과]
	5 * 9 = 45
*/
