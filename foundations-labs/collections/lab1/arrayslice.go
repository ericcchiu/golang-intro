package lab1

/*BasicArray Exercise 1:
Create an array named 'numbers' that contains, in order, 1,2,3,4, and 5 */
func BasicArray() [5]int {
	//=========================Code Below =================================//
	numbers := [5]int{1, 2, 3, 4, 5}
	//=========================Code Above =================================//
	return numbers
}

/*BasicSlice Exercise 2: Create a basic Slice
Create a slice named 'food' that contains "apple", "pizza", "taco", "cheese"
*/
func BasicSlice() []string {
	//=========================Code Below =================================//
	food := []string{"apple", "pizza", "taco", "cheese"}
	//=========================Code Above =================================//
	return food
}

/*MergeAndSubSlice Exercise 3:
  Merge a slice with another and get a sub slice
  STEP 1: Add the morefood slice to the food slice
  STEP 2:  Get the subslice between index 1 and 3 and assign it to a variable named sublist
  Hint: variables sublist & morefood are already defined for you outside the codeblock
*/
func MergeAndSubSlice(morefood []string) (food []string, sublist []string) {

	food = []string{"apple", "pizza"}
	//=========================Code Below =================================//
	food = append(food, morefood...)
	sublist = morefood[1:4]
	//=========================Code Above =================================//
	return
}
