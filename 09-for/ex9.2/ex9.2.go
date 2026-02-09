package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	for {
		time.Sleep(time.Second) // 쓰레드를 1초동안 중지
		fmt.Println(i)
		i++
	}
}
