package arrays

// RemoveDuplicates function removes array duplicates in-place such that each
// element appears only onceand return the new length
func RemoveDuplicates(nums []int) int {
	// Create a map that records duplicates we find them
	mapOccurrence := map[int]bool{}
	array := []int{}

	for val := range nums {
		if mapOccurrence[nums[val]] == true {
			// Do not add duplicate
		} else {
			mapOccurrence[nums[val]] = true
			array = append(array, nums[val])
		}
	}

	return len(array)
}
