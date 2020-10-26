package conditionals1

import "fmt"

func isOldEnoughToDrink(age int) bool {
	if age >= 21 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isOldEnoughToDrink(22))
}
