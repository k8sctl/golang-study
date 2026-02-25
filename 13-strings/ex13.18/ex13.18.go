package main

import (
	"fmt"
	"strings"
)

// ToUpper1: 문자열 덧붙이기(+=)로 직접 대문자 변환
// 동작은 맞지만, += 는 누적할수록 복사가 반복되어 비효율적일 수 있음
func ToUpper1(str string) string {
	var rst string
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			rst += string('A' + (c - 'a'))
		} else {
			rst += string(c)
		}
	}
	return rst
}

// ToUpper2: strings.Builder로 대문자 변환(권장)
// Builder에 rune을 쌓고 마지막에 String()으로 한 번에 결과 생성
func ToUpper2(str string) string {
	var builder strings.Builder
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a'))
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func main() {
	var str string = "Hello World"

	fmt.Println(ToUpper1(str))
	fmt.Println(ToUpper2(str))
}

/*
[코드 설명]
- 이 코드는 문자열을 "대문자"로 바꾸는 함수를 2가지 방식으로 구현하고,
  같은 입력("Hello World")에 대해 결과를 비교 출력한다.

1) ToUpper1(str string) string
- rst라는 문자열을 빈 문자열로 시작한다.
- for-range로 str을 "rune(유니코드 문자)" 단위로 순회한다.
- 문자가 'a'~'z'이면:
  - 'A' + (c - 'a')로 대문자 알파벳으로 변환한다.
    예) 'c' - 'a' = 2, 'A' + 2 = 'C'
- 변환된 문자를 string(...)으로 바꾼 뒤 rst += 로 계속 이어 붙인다.
- 주의: rst += 는 매번 새 문자열을 만들어 복사하는 비용이 들어서,
  문자열이 길어질수록 비효율적일 수 있다.

2) ToUpper2(str string) string
- strings.Builder를 사용해 문자를 누적한다.
- for-range로 str을 rune 단위로 순회하며,
  'a'~'z'는 대문자로 변환해서 builder.WriteRune(...)로 추가하고,
  그 외 문자는 그대로 builder.WriteRune(c)로 추가한다.
- 마지막에 builder.String()으로 누적된 결과를 한 번에 문자열로 만든다.
- 일반적으로 Builder 방식이 문자열 누적(+=)보다 효율적이다.

3) main()
- str := "Hello World"를 준비한다.
- ToUpper1(str), ToUpper2(str)의 결과를 각각 출력한다.
- 출력은 둘 다 동일하게 "HELLO WORLD"가 된다.
*/

/*
[출력 결과]
HELLO WORLD
HELLO WORLD
*/
