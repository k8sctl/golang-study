package custompkg

import "fmt"

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
