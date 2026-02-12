package main

import (
	"fmt"
)

// 배열 순회 2
/*
range로 배열의 인덱스(i) 와 값(v) 을 동시에 꺼내서 출력하는 예제
- i: 0부터 시작하는 인덱스
- v: 해당 인덱스의 값 (float64)
- 배열 t 길이가 5니까 총 5번 반복한다.
*/

func main() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	for i, v := range t {
		fmt.Println(i, v)
	}
}

/*
[출력 결과]
0 24
1 25.9
2 27.8
3 26.9
4 26.2
*/
