// https://leetcode.com/problems/running-sum-of-1d-array/

// Test cases:
// nums = [1,2,3,4]
// nums = [1,1,1,1,1]
// nums = [3,1,2,10,1]

// Expected results:
// [1,3,6,10]
// [1,2,3,4,5]
// [3,4,6,16,17]

package leetcode

// for cycle solution
func runningSum1(nums []int) []int {
	sum := 0
	result := make([]int, len(nums)) // create an empty array of integers with the same length as nums

	for i, e := range nums {
		sum += e
		result[i] = sum
	}

	return result
}

// Override initial array solution
func runningSum2(nums []int) []int {
	for i := range nums {
		if i == 0 {
			continue
		}

		nums[i] += nums[i-1]
	}

	return nums
}
