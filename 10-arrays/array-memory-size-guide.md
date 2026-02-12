## 배열 크기 정리

배열이 차지하는 메모리 크기는 기본적으로 아래처럼 계산한다.

- 1차원 배열: `배열 크기 = 요소 타입 크기 × 요소 개수`
- 다차원 배열: `배열 크기 = 요소 타입 크기 × 전체 요소 개수(각 차원 길이의 곱)`

예시 계산:
- `[10]int32` → `4바이트 × 10 = 40바이트`
- `[2][5]int32` → `4바이트 × (2×5) = 40바이트`
- `[3][4]float64` → `8바이트 × (3×4) = 96바이트`

Go의 다차원 배열은 배열 안에 배열 형태이며, 고정 길이 배열이므로 메모리에 연속적으로 배치된다.  
따라서 1차원/2차원 모두 주소 계산 기반으로 빠르게 접근할 수 있다.

다만 타입 관점에서는 `[10]int`와 `[2][5]int`는 서로 다른 타입이다.

---

## 예시 코드 (unsafe.Sizeof로 배열 크기 확인)

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a [10]int32
	var b [2][5]int32
	var c [3][4]float64

	fmt.Println("sizeof(a):", unsafe.Sizeof(a)) // 40
	fmt.Println("sizeof(b):", unsafe.Sizeof(b)) // 40
	fmt.Println("sizeof(c):", unsafe.Sizeof(c)) // 96
}