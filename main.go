package main

import (
	"fmt"
)

func main() {

	y := 4
	if x := y % 2; x == 0 {
		fmt.Println("Even")
	}
	fmt.Println("Odd")
	// fmt.Println(x)
}
