package custompkg

import "fmt"

// Student: 대문자 시작 → export 됨 (패키지 외부에서 custompkg.Student로 사용 가능)
type Student struct {
	// Name: 대문자 시작 → export 됨 (외부에서 student.Name 접근/수정 가능)
	Name string

	// Age: 대문자 시작 → export 됨 (외부에서 student.Age 접근/수정 가능)
	Age int

	// score: 소문자 시작 → unexported (패키지 외부에서 접근/수정 불가)
	// 같은 패키지(custompkg) 내부 코드에서만 접근 가능
	score int
}

// Var1: 대문자 시작 → export 됨 (패키지 외부에서 custompkg.Var1로 접근 가능)
var Var1 int

// var2: 소문자 시작 → unexported (패키지 외부에서 custompkg.var2 접근 불가)
var var2 int

// PI: 대문자 시작 → export 됨 (패키지 외부에서 custompkg.PI로 접근 가능)
const PI = 3.14

// pI2: 소문자 시작 → unexported (패키지 외부에서 custompkg.pI2 접근 불가)
const pI2 = 3.1415

// PrintCustom: 대문자 시작 → export 됨
func PrintCustom() {
	fmt.Println("This is custom package")
}

// printCustom2는 export되지 않는(패키지 내부 전용) 함수다.
func printCustom2() {
	fmt.Println("This is custom package222!")
}

/*
[코드 설명]
- custompkg 패키지는 로컬 패키지 예시다.

1) PrintCustom()
- 대문자 P로 시작 → export 됨
- 다른 패키지(main 등)에서 custompkg.PrintCustom() 형태로 호출 가능

2) printCustom2()
- 소문자 p로 시작 → export 안 됨
- custompkg 내부에서만 호출 가능
- main 패키지에서 custompkg.printCustom2()로 호출하면 컴파일 에러가 발생

[출력 결과]
This is custom package
(단, printCustom2()는 외부에서 호출되지 않으면 출력되지 않음)
*/
