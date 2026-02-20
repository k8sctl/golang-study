package main

import "fmt"

func main() {
	str1 := "Hello"
	str2 := "World"

	str3 := str1 + " " + str2
	fmt.Println(str3)

	str1 += " " + str2
	fmt.Println(str1)
}

/*
[코드 설명]
- 문자열은 + 연산자로 이어 붙일 수 있다(concatenation).
- str3 := str1 + " " + str2
  -> "Hello", 공백, "World"를 합쳐서 "Hello World"를 만든다.
- str1 += " " + str2
  -> str1에 " World"를 추가로 붙여서 str1 자체가 "Hello World"가 된다.
  (Go 문자열은 불변(immutable)이라 내부적으로는 새 문자열이 만들어지고 str1이 그 값을 참조하게 된다.)
*/

/*
[출력 결과]
Hello World
Hello World
*/
