package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(200)
	fmt.Println("x is %i", x)
	// Insert code below
	switch {
	case x >= 1 && x <= 25:
		fmt.Println("First Quarter")
	case x >= 26 && x <= 50:
		fmt.Println("Second Quarter")
	case x >= 51 && x <= 100:
		fmt.Println("Second Half")
	default:
		fmt.Println("Out of Range")
	}
}
