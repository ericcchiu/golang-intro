package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(200)
	// Insert code below this line
	fmt.Println("x is %i", x)
	if x >= 1 && x <= 25 {
		fmt.Println("First Quarter")
	} else if x >= 26 && x <= 50 {
		fmt.Println("Second Quarter")
	} else if x >= 51 && x <= 100 {
		fmt.Println("Second Half")
	} else {
		fmt.Println("Out of Range")
	}
}
