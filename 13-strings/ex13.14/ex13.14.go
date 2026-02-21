package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	str1 := "Hello 월드"
	str2 := str1

	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	fmt.Println(stringHeader1)
	fmt.Println(stringHeader2)
}

/*
[출력 결과]
&{4963388 12}
&{4963388 12}

---

&{<주소값> 12}
&{<주소값> 12}
*/

/*
[코드 설명]
- 이 예제는 Go 문자열(string)이 내부적으로
  (1) 문자열 데이터가 있는 메모리 주소(Data)
  (2) 문자열 길이(바이트 단위, Len)
  로 이루어진 '헤더(header)' 형태라는 것을 확인하는 코드다.

- str2 := str1 은 문자열 내용을 새로 복사하기보다는,
  보통 str1이 가리키는 문자열 데이터 주소와 길이 정보를 그대로 복사한다.
  (문자열은 불변(immutable)이므로 같은 데이터를 공유해도 안전하다.)

[reflect.StringHeader + unsafe 사용]
- stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
  - &str1  : string 변수(str1) 자체의 주소
  - unsafe.Pointer(...) : 타입 안전 검사를 무시하고 포인터 변환
  - (*reflect.StringHeader)(...) : 그 메모리를 StringHeader 구조체처럼 “해석”해서
    내부 필드(Data, Len)를 읽어온다.

- stringHeader2도 str2에 대해 같은 방식으로 헤더를 읽는다.

[출력 해석]
- 출력 형태: &{Data Len}
  예) &{4963388 12}
  - Data: 문자열 데이터 시작 주소(실행마다 달라질 수 있음)
  - Len : 문자열 길이(바이트 단위)

- "Hello 월드"의 Len이 12인 이유(UTF-8 기준):
  "Hello "(6바이트) + "월"(3바이트) + "드"(3바이트) = 12바이트

- 대부분의 경우 str1과 str2는 같은 문자열 데이터를 공유하므로
  Data 값이 같게 출력된다.
  즉, 두 문자열이 같은 메모리의 문자열 데이터를 가리킨다는 의미다.

[주의]
- 이 코드는 학습용 확인 목적에 적합하지만,
  reflect.StringHeader를 unsafe로 캐스팅하는 방식은 최신 Go에서 권장되지 않는다.
  (경고가 뜰 수 있음)
*/
