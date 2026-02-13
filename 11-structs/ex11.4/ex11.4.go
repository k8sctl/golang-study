package main

import (
	"fmt"
)

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	User     // 내장된 필드(embedded field) 방식으로 선언
	VIPLevel int
	Price    int
	// Name string
	// 만약 VIPUser에 Name 필드가 있다면 내장된 필드보다 VIPUser에 선언된 Name 필드가 우선적으로 적용된다.
	// 만약 내장된 필드를 사용하고 싶으면 vip.User.Name으로 접근하면 해당 값에 접근할 수 있다.
}

func main() {
	user := User{"송하나", "hana", 23}
	vip := VIPUser{
		User{"화랑", "hwarang", 48},
		3,
		250,
	}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP레벨: %d 가격: %d만원\n",
		vip.Name,
		vip.ID,
		vip.Age,
		vip.VIPLevel,
		vip.Price,
	)
}

/*
[출력 결과]
유저: 송하나 ID: hana 나이: 23
VIP 유저: 화랑 ID: hwarang 나이: 48 VIP레벨: 3 가격: 250만원
*/
