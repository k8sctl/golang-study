package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		[입력 버퍼]
		표준 입력과 표준 출력은 비트스트림 형태로 입력되고 출력된다.
		이렇게 비트 스트림 형태로 들어오는 데이터가 잠시 들어오는 임시 저장소를 버퍼라고 한다.

		입력 실패를 했을 때, 기존에 실패했던 값들이 지워지지 않고 입력 버퍼에 남아있기 때문에 버퍼를 비우는 작업이 필요하다.
	*/

	stdin := bufio.NewReader(os.Stdin) // var stdin = bufio.NewReader(os.Stdin)과 같다.

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
		// 특정 문자열(지금은 개행 문자)이 나올 때까지 읽어라!
		// 이 부분을 통해 err를 출력할 때, 입력 버퍼를 비울 수 있다.
	} else {
		fmt.Println(n, a, b) // err가 발생하지 않았다면 정상적인 상황이기 때문에 a, b를 출력하면 된다.
	}

	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}

	/*
		hello 4 		 // 첫 번째 입력
		expected integer // err 출력
		10 4			 // 두 번째 입력
		2 10 4			 // 정상적인 출력, 여기서 n은 입력 받은 값의 개수를 출력하기 때문에 2를 출력한다.
	*/
}
