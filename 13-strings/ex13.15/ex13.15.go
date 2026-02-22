package main

import "fmt"

func main() {
	var str string = "Hello World"
	var slice []byte = []byte(str) // string -> []byte 변환(복사본 생성)

	slice[2] = 'a' // slice는 수정 가능하므로 3번째 바이트를 'a'로 변경

	fmt.Println(str)          // 원본 문자열은 그대로
	fmt.Printf("%s\n", slice) // 수정된 결과 출력
}

/*
[코드 설명: 문자열 불변(immutable)과 []byte 변환]

- Go의 string은 불변(immutable) 타입이다.
  => 문자열의 "일부 바이트"를 직접 수정할 수 없다.
     예: str[2] = 'a'  // 컴파일 에러

- 하지만 문자열 변수에 "새 문자열"을 통째로 다시 대입하는 것은 가능하다.
  예:
    str = "How are you?"  // OK (새 문자열로 재할당)

- 문자열 일부를 바꾸고 싶다면, 수정 가능한 타입으로 변환해서 작업해야 한다.
  대표적으로 []byte 또는 []rune를 사용한다.
  (ASCII 중심이면 []byte, 유니코드 문자 단위면 []rune가 안전)

[코드의 흐름]
1) str := "Hello World"
2) slice := []byte(str)
   - 문자열을 바이트 슬라이스로 "복사"해서 만든다.
   - slice는 수정 가능(mutable)하다.
3) slice[2] = 'a'
   - 0부터 시작하는 인덱스 기준 2번째(3번째) 문자를 수정한다.
   - "Hello World"에서 'l'이 'a'로 바뀌어 "Healo World"가 된다.
4) fmt.Println(str)
   - 원본 string은 불변이며, slice는 복사본이므로 str은 그대로 출력된다.
5) fmt.Printf("%s\n", slice)
   - 수정된 바이트 슬라이스를 문자열처럼 출력한다.
*/

/*
[출력 결과]
Hello World
Healo World
*/
