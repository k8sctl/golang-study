package main

import "fmt"

func getMyAge() int {
	return 28
}

func main() {
	switch age := getMyAge(); age {
	case 10:
		fmt.Println("Teenage")
	case 33:
		fmt.Println("Pair 3")
	default:
		fmt.Println("My age is", age)
	}

	// fmt.Println("age is", age) // age 변수는 소멸했기 때문에 에러가 발생한다.
}
