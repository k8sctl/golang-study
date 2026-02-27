package main

import (
	"fmt"
	"golang/usepkg/custompkg"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main() {
	custompkg.PrintCustom()
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}

/*
[코드 설명]
- 이 코드는 “패키지 import + 모듈 의존성 관리(go.mod)”를 한 번에 보여주는 예제다.

1) import 구분
- "fmt" : Go 표준 라이브러리(기본 제공)
- "golang/usepkg/custompkg" : 네 프로젝트 내부(로컬) 패키지
- "github.com/guptarohit/asciigraph" : 외부(서드파티) 모듈 패키지
- "github.com/tuckersGo/musthaveGo/ch16/expkg" : 외부(서드파티) 모듈 패키지

2) main 흐름
- custompkg.PrintCustom(): 로컬 패키지 함수 호출
- expkg.PrintSample(): 외부 패키지 함수 호출
- asciigraph.Plot(data): float64 배열을 ASCII 그래프로 변환해서 출력

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
