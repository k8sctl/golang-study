package main

import (
	"fmt"
)

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

func main() {
	/*
		[쇼트 서킷 평가(Short-Circuit Evaluation)]
		- && 연산에서 좌변이 false 이면, 결과는 이미 false로 결정되므로 우변을 평가하지 않는다.
		- || 연산에서 좌변이 true 이면, 결과는 이미 true로 결정되므로 우변을 평가하지 않는다.

		[주의할 점]
		우변에 함수 호출이 있으면, 쇼트 서킷으로 인해 함수가 실행되지 않을 수 있다.
		즉, 우변 함수에 로그 기록/상태 변경 같은 부수효과(side effect) 가 있다면 의도와 다른 동작이 발생할 수 있다.
	*/

	if false && IncreaseAndReturn() < 5 {
		// 위에서 설명한 것처럼, && 연산에서 좌변이 false이기 때문에 IncreaseAndReturn() 함수는 호출되지도 않는다.
		fmt.Println("1 증가")
	}

	/*
		이러한 이유르 if 조건문에 들어가는 함수 호출는
		값을 조작하고 연산하는 것보다 단순 값만 비교하는 것이 좋다.
		(연산이 일어나지 않을 수도 있기 때문에)
	*/
}
