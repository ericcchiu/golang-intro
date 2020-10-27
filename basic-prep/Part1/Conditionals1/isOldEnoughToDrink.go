package conditionals1

// IsOldEnoughToDrink function takes in a person's age (integer) and determines if they are 21 or older with output of true or false.
func IsOldEnoughToDrink(age int) bool {
	if age >= 21 {
		return true
	}
	return false
}
