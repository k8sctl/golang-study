package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var str string = "Hello"
	addr1 := unsafe.StringData(str)
	str += " World"
	addr2 := unsafe.StringData(str)
	str += " Welcome!"
	addr3 := unsafe.StringData(str)
	fmt.Println(str)
	fmt.Printf("addr1:\t%p\n", addr1)
	fmt.Printf("addr2:\t%p\n", addr2)
	fmt.Printf("addr3:\t%p\n", addr3)
}

/*
[코드 설명]
- Go의 string은 불변(immutable)이라서 `str += "..."` 같은 연산은
  기존 문자열을 수정하지 않고, "새 문자열"을 만들어 str에 다시 대입한다.
  (즉, 결합이 일어날 때마다 새로운 문자열 데이터가 만들어진다.)

- unsafe.StringData(str)는 str이 가리키는 문자열 데이터의 시작 주소(*byte)를 반환한다.
  addr1/addr2/addr3를 비교하면, 문자열 결합 전/후에
  문자열 데이터의 시작 주소가 어떻게 달라지는지 관찰할 수 있다.

- 흐름:
  1) str = "Hello"                   -> addr1
  2) str += " World"                 -> str = "Hello World"            -> addr2
  3) str += " Welcome!"              -> str = "Hello World Welcome!"   -> addr3
  4) 최종 문자열과 각 시점의 데이터 주소를 출력한다.

[주의]
- string 결합은 새 문자열을 만들지만, 출력되는 주소 값 자체는 실행 환경/런타임 상태에 따라 달라질 수 있다.
- unsafe 패키지는 학습/검증 목적 외에는 사용을 지양하는 것이 좋다.
*/

/*
[출력 결과]
Hello World Welcome!
addr1:  0x....
addr2:  0x....
addr3:  0x....

- addr1/addr2/addr3는 보통 서로 다른 값으로 출력된다(각 결합 결과가 새 문자열이기 때문).
- 주소 값은 실행할 때마다 달라질 수 있다.
*/
