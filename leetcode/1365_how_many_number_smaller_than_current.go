package goleetcode

import (
	"sort"
)

// Learning: Sorting int slice, debug logging.
func SmallerNumbersThanCurrent(nums []int) []int {
	if len(nums) == 0 {
		return make([]int, 0)
	}

	// Count smaller numbers.
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	smallerNumbersCount := make(map[int]int)
	smallerNumbersCount[sortedNums[0]] = 0
	for i := 1; i < len(sortedNums); i++ {
		if sortedNums[i] != sortedNums[i-1] {
			smallerNumbersCount[sortedNums[i]] = i
			// fmt.Println(sortedNums[i], i)
		}
	}

	// Populate result slice.
	result := make([]int, len(nums))
	for i := range nums {
		result[i] = smallerNumbersCount[nums[i]]
	}

	return result
}
