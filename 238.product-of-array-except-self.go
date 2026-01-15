package leetcode

// @leet start
func productExceptSelf(nums []int) []int {
	sumProducts := make([]int, len(nums))

	for i := range nums {
		if i == 0 {
			sumProducts[i] = 1
			continue
		}

		sumProducts[i] = sumProducts[i-1] * nums[i-1]
	}

	right := 1
	for k := len(nums) - 1; k >= 0; k-- {
		sumProducts[k] *= right
		right *= nums[k]
	}

	return sumProducts
}

// @leet end
