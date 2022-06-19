package main

import (
	"fmt"
)

func main() {
	// bit shifting
	a := 8 // 2^3
	// left
	fmt.Println(a << 3) // 2^3 * 2^3 = 2^6
	// right
	fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0
}
