package main

import "fmt"

/*
	[중첩 if 구문]
	음식값이 5만원이 넘고 친구 중에 부자가 있다면 신발끈을 묶는다. 부자가 없다면 돈을 나눠 낸다.
	음식값이 3만원 이상이고 5만원 이하이고 같이 간 친구가 3명이 넘으면 신발끈을 묶는다. 3명 이하면 돈을 나눠낸다.
	3만원 미만이면 내가 낸다.
*/

func HasRichFriend() bool {
	return true
}

func GetFriendsCount() int {
	return 3
}

func main() {
	price := 35000 // 음식값이 35000원

	if price >= 50000 {
		if HasRichFriend() {
			fmt.Println("앗 신발끈이 풀렸네!")
		} else {
			fmt.Println("나눠내자.")
		}
	} else if price >= 30000 {
		if GetFriendsCount() > 3 {
			fmt.Println("어이쿠 신발끈이...")
		} else {
			fmt.Println("사람도 얼마 없는데 나눠내자.") // 출력
		}
	} else {
		fmt.Println("오늘은 내가 쏜다!")
	}
}
