package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ { // 바깥 루프: 3번 반복(행)
		for j := 0; j < 5; j++ { // 안쪽 루프: 매번 5번 반복(열)
			fmt.Print('*') // * 을 출력
		}
		fmt.Println() // 개행한다.
	}
}

/*
	[출력 결과]
	*****
	*****
	*****
*/
