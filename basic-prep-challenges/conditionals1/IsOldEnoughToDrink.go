/*
Write a function called "isOldEnoughToDrink". Given a number, in this case an age, "isOldEnoughToDrink" returns whether a person of this given age is old enough to legally drink in the United States. Notes:

The legal drinking age in the United States is 21.
*/
package conditionals1

func IsOldEnoughToDrink(age int) bool {
	if age >= 21 {
		return true
	}
	return false
}
