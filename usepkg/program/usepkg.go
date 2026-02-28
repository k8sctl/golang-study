package main

import (
	"fmt"

	"golang/usepkg/custompkg"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main() {
	custompkg.PrintCustom()
	// custompkg.printCustom2() // 소문자 시작: 다른 패키지에서 접근 불가(비공개)
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}

/*
[코드 설명]
- 이 코드는 “패키지 import 방식”과 “모듈 의존성(go.mod)로 외부 패키지 가져오기”를 함께 보여주는 예제다.

1) import 대상 구분
- "fmt"
  - Go 표준 라이브러리(기본 제공)

- "golang/usepkg/custompkg"
  - 현재 모듈 내부의 로컬 패키지 import
  - go.mod의 module 경로가 "golang/usepkg" 라고 가정하면,
    "golang/usepkg/custompkg"는 custompkg 디렉토리를 가리킨다.

- "github.com/guptarohit/asciigraph"
  - 외부 모듈(서드파티) 패키지
  - go get / go mod tidy 등을 통해 go.mod에 의존성이 기록된다.

- "github.com/tuckersGo/musthaveGo/ch16/expkg"
  - 외부 모듈(서드파티) 패키지

2) main 함수 흐름
- custompkg.PrintCustom()
  - 로컬 패키지(custompkg)의 공개 함수 호출(대문자 시작이라 export 됨)

- custompkg.printCustom2() 호출이 주석 처리된 이유
  - 소문자 시작 함수는 export되지 않아서(패키지 외부에서 접근 불가) 컴파일 에러가 난다.

- expkg.PrintSample()
  - 외부 패키지 함수 호출

- asciigraph.Plot(data)
  - float64 슬라이스를 ASCII 그래프로 렌더링해 문자열로 반환
  - fmt.Println으로 출력

[출력 결과]
This is custom package
This is Github expkg Sample
 10.00 ┤        ╭╮
  9.00 ┤   ╭╮   ││
  8.00 ┤   ││ ╭╮││
  7.00 ┤   │╰╮││││╭╮
  6.00 ┤  ╭╯ │││││││ ╭
  5.00 ┤ ╭╯  ╰╯╰╯│││╭╯
  4.00 ┤╭╯       ││││
  3.00 ┼╯        ││││
*/
