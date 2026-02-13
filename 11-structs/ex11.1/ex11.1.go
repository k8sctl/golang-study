package main

import (
	"fmt"
)

type House struct {
	Address  string
	Size     int
	Price    float64
	Category string
}

func main() {
	var house House
	house.Address = "서울시 강남구..."
	house.Size = 28
	house.Price = 10
	house.Category = "아파트"

	fmt.Println(house) // fmt.Printf("%v\n", house)와 같다.
	fmt.Printf("%v\n", house)
	fmt.Printf("주소:%s 사이즈:%d평 가격:%f억원 종류:%s\n", house.Address, house.Size, house.Price, house.Category)
}

/*
[출력 결과]
{서울시 강남구... 28 10 아파트}
{서울시 강남구... 28 10 아파트}
주소:서울시 강남구... 사이즈:28평 가격:10.0억원 종류:아파트
*/
