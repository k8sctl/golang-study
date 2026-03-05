package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		[랜덤 시드 설정]
		- rand.Seed(seed)는 math/rand의 "전역" 난수 생성기의 시드를 설정한다.
		- time.Now()로 현재 시각(Time)을 얻고,
		  t.UnixNano()로 Unix epoch(1970-01-01 UTC)부터 경과한 나노초(int64)를 구해 시드로 사용한다.
		- 이렇게 하면 프로그램을 실행할 때마다 시드가 달라져 난수 결과도 달라질 가능성이 크다.

		[참고: rand.Seed 관련 경고(Go 1.20+)]
		- 일부 IDE/정적 분석 도구(staticcheck/go vet 등)에서 아래 취지의 경고가 뜰 수 있다:
		  "Go 1.20부터는 math/rand의 전역 난수 생성기가 기본적으로 랜덤 시드로 초기화되므로,
		   실행마다 랜덤을 원한다면 rand.Seed(time.Now().UnixNano()) 호출은 보통 불필요하다."
		- 즉 Go 1.20+ 환경에서는 대개 아래처럼만 써도 실행마다 다른 난수가 나온다:
		  rand.Intn(100)

		[고정 시드(재현 가능한 결과)가 필요할 때 권장 방식]
		- 테스트/디버깅처럼 항상 같은 난수 시퀀스가 필요하면
		  전역 rand.Seed(...) 대신 "로컬 RNG"를 만들어 사용한다:
		  r := rand.New(rand.NewSource(1)) // 고정 시드
		  n := r.Intn(100)

		[난수 생성]
		- rand.Intn(100)은 0 ~ 99 범위의 난수를 반환한다.
		- 아래 코드는 0~99 난수를 10개 출력한다.
	*/
	t := time.Now()
	rand.Seed(t.UnixNano())

	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
	}
}

/*
[출력 결과]
(실행할 때마다 달라짐)

예시)
42
7
98
13
0
64
21
88
5
73
*/
