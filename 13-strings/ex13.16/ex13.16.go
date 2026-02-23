package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
[코드 설명]
- 이 코드는 string과 []byte가 “같은 데이터를 공유하는지” 확인하기 위해
  내부 데이터 시작 주소(Data)를 출력하는 예제다.

- str := "Hello World"
  -> 문자열(string)은 내부적으로 (Data 포인터, Len) 형태의 헤더를 가진다.

- slice := []byte(str)
  -> string을 []byte로 변환하면 보통 “복사(copy)”가 일어난다.
     즉, slice는 별도의 바이트 배열을 새로 만들고 그 값을 복사해 가진다.
     (그래서 slice를 수정해도 원본 str이 바뀌지 않는다.)

- reflect.StringHeader / reflect.SliceHeader + unsafe를 이용해
  각각의 Data(실제 데이터가 시작되는 메모리 주소)를 읽어와 출력한다.

- 출력된 str.Data 와 Slice.Data 가 다르면:
  => string과 []byte가 서로 다른 메모리를 사용(복사됨)
- 만약 같다면:
  => 같은 메모리를 공유(일반적인 []byte(str) 변환에서는 거의 기대하지 않음)

[주의]
- reflect.StringHeader / SliceHeader를 unsafe로 캐스팅하는 방식은
  최신 Go에서 권장되지 않을 수 있으며(Deprecated 경고),
  학습/확인 목적으로만 사용하는 것이 안전하다.
*/

func main() {
	var str string = "Hello World"
	var slice []byte = []byte(str)

	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))

	fmt.Printf("str:\t%x\n", stringHeader.Data)  // str이 가리키는 문자열 데이터 주소
	fmt.Printf("slice:\t%x\n", sliceHeader.Data) // slice의 내부 배열 시작 주소
}

/*
[출력 결과]
str:    <주소값1>
slice:  <주소값2>

- 보통 <주소값1> 과 <주소값2> 는 다르게 출력된다.
  왜냐하면 []byte(str) 변환 시 문자열 데이터를 새 바이트 배열로 복사하기 때문이다.
*/
