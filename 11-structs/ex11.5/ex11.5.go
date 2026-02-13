package main

import (
	"fmt"
)

type Student struct {
	Age   int // 대문자로 시작하는 필드는 외부로 공개된다.
	No    int
	Score float64
}

func PrintStudent(s Student) {
	fmt.Printf("나이: %d 번호: %d 점수:%.2f\n",
		s.Age,
		s.No,
		s.Score,
	)
}

func main() {
	student := Student{15, 23, 88.2}

	student2 := student // student 구조체의 모든 필드가 student2로 복사된다.

	PrintStudent(student)
	PrintStudent(student2)
}

/*
[출력 결과]
나이: 15 번호: 23 점수:88.20
나이: 15 번호: 23 점수:88.20
*/
