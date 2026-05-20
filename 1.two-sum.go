package leetcode

// @leet start
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))

	for i, num := range nums {
		diff := target - num
		if pos, ok := seen[diff]; ok {
			return []int{pos, i}
		}

		seen[num] = i

	}

	return []int{}
}

// @leet end

