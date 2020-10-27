package conditionals1

// IsOldEnoughToDrive function takes in an integer argument age and determines if person is old enough to drive at age 16 or greater.
func IsOldEnoughToDrive(age int) bool {
	if age >= 16 {
		return true
	}
	return false
}
