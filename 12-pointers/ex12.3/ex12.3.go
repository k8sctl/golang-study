package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func changeData(arg *Data) {
	arg.value = 999
	arg.data[100] = 999 // 배열의 101번째 요소
}

func main() {
	var data Data

	changeData(&data)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
}

/*
[출력 결과1]
value = 0
data[100] = 0

[코드 설명1]
- Data는 value(int)와 data([200]int)를 가지는 구조체다.
- changeData(arg Data)는 구조체를 "값으로" 받는다.
  즉, main의 data가 그대로 전달되는 것이 아니라 "복사본"이 전달된다.
- 함수 내부에서 arg.value, arg.data[100]을 바꿔도 복사본만 변경된다.
- 함수가 끝나면 복사본은 사라지고, 원본 data는 그대로 유지된다.
- 그래서 출력은 value=0, data[100]=0 이다.

[핵심]
- Go에서 구조체/배열은 기본적으로 값 타입이다.
- 원본을 바꾸려면 포인터(예: *Data)를 사용해야 한다.

---

[출력 결과2]
value = 999
data[100] = 999

[코드 설명2]
- Data는 value(int)와 data([200]int)를 가지는 구조체다.
- changeData(arg *Data)는 구조체 포인터를 매개변수로 받는다.
- changeData(&data)로 원본 data의 주소를 전달했기 때문에,
  함수 내부에서 arg.value, arg.data[100]을 변경하면 원본 data가 직접 변경된다.
- 따라서 함수 호출 후 출력값은 value=999, data[100]=999 가 된다.

[핵심]
- 값 전달(arg Data): 복사본 변경(원본 불변)
- 포인터 전달(arg *Data): 원본 직접 변경
*/
