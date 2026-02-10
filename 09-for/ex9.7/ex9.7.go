package main

import "fmt"

/*
[중첩 for문 종료 방법]
중첩 for문에서 특정 조건이 만족되면 바깥 반복문까지 종료해야 할 때는 다음 방법을 사용한다.

1) 플래그 변수 사용
   - 종료 여부를 bool 변수로 저장하고,
   - 안쪽 루프에서 값을 바꾼 뒤 바깥 루프에서 확인하여 종료한다.
*/

func main() {
	a := 1
	b := 1
	found := false
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}

/*
	[출력 결과]
	5 * 9 = 45
*/
