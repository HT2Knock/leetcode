package leetcode

// @leet start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nextGreater := make(map[int]int, len(nums1))
	stack := make([]int, 0, len(nums2))

	for _, val := range nums2 {
		for len(stack) > 0 && val > stack[len(stack)-1] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			nextGreater[top] = val
		}

		stack = append(stack, val)
	}

	result := make([]int, len(nums1))
	for i, num := range nums1 {
		if val, exist := nextGreater[num]; exist {
			result[i] = val
			continue
		}

		result[i] = -1
	}

	return result
}

// @leet end
