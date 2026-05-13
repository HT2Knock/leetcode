package leetcode

// @leet start
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	mid := (left + right) / 2

	for left <= right {
		mid = (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		}
	}

	return left
}

// @leet end

