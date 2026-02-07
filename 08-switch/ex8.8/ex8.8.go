package main

import "fmt"

type ColorType int // 별칭 타입
const (
	Red    ColorType = iota // 0
	Blue                    // 1
	Green                   // 2
	Yellow                  // 3
)

func ColorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}

func getMyFavoriteColor() ColorType {
	return Red
}

func main() {
	fmt.Println("My favorite color is", ColorToString(getMyFavoriteColor()))
	fmt.Println("My favorite color is", ColorToString(1))
	fmt.Println("My favorite color is", ColorToString(Blue))
}
