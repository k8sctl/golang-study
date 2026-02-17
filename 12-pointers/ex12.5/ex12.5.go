package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	var u = User{name, age}
	return &u
}

func main() {
	userPointer := NewUser("AAA", 23)
	fmt.Println(userPointer)
}

/*
[출력 결과]
&{AAA 23}
*/

/*
[코드 설명]
- User 구조체는 Name(string), Age(int) 필드를 가진다.
- NewUser(name, age) 함수는 User 값을 만든 뒤 그 주소(*User)를 반환한다.
  즉, 구조체 생성자처럼 동작하는 함수다.
- main에서 userPointer는 *User 타입(구조체 포인터)이다.
- fmt.Println(userPointer)는 포인터가 가리키는 구조체 값을 '&{필드값...}' 형태로 출력한다.
  그래서 결과가 &{AAA 23} 으로 나온다.

[핵심]
- 값(User)이 아니라 포인터(*User)를 반환하면,
  같은 인스턴스를 공유하거나 함수/메서드에서 원본 데이터를 직접 수정하기 쉽다.
*/

/*
[질문] 함수 내부 지역 변수(u)의 주소를 반환해도 괜찮을까?
- NewUser() 안의 u는 지역 변수이므로, 함수가 끝나면 사라지는 것처럼 보인다.
- 그런데도 &u를 반환해도 Go에서는 안전하다.

[이유: Escape Analysis]
- 컴파일러가 "u의 주소가 함수 밖으로 나간다(escape)"는 것을 분석한다.
- 이런 경우 u를 스택이 아니라 힙(Heap)에 할당해 함수 종료 후에도 유효하게 유지한다.
- 따라서 반환된 포인터(*User)는 유효한 인스턴스를 가리킨다.

[언제 메모리에서 사라지나?]
- 반환된 포인터를 더 이상 어디에서도 참조하지 않으면
  해당 인스턴스는 도달 불가능(unreachable) 상태가 된다.
- 이후 가비지 컬렉터(GC)가 실행될 때 메모리에서 회수된다.

[정리]
- "사라진 지역 변수의 주소를 반환"하는 것이 아니라,
  컴파일러가 필요 시 힙으로 옮겨 안전하게 관리해 주는 것이다.
*/
