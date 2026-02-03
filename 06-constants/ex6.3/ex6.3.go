package main

import (
	"fmt"
)

// 코드값으로 사용
// 코드란 숫자에 의미를 부여하는 것이다.
// 대표적인 예시로 ASCII 코드가 존재한다.
// 'A'=65, 'B'=66, 'C'=67, 'D'=68 ...

// UNICODE 2byte로 문자를 표현하는 문자코드
// UTF-8 1~3byte로 문자를 표현하는 문자코드
// ANSI(ASCII) 1byte로 문자를 표현하는 문자코드

const Pig int = 0
const Cow int = 1
const Chicken int = 2

func PrintAnimal(animal int) {
	if animal == Pig {
		fmt.Println("꿀꿀")
	} else if animal == Cow {
		fmt.Println("음메")
	} else if animal == Chicken {
		fmt.Println("꼬끼오")
	} else {
		fmt.Println("...")
	}
}

func main() {
	PrintAnimal(Cow)
	PrintAnimal(Pig)
	PrintAnimal(7)
}
