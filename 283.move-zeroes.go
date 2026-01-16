package leetcode

// @leet start
func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}

	anchor := 0

	for explorer := 0; explorer < len(nums); explorer++ {
		if nums[explorer] != 0 {
			if explorer != anchor {
				nums[anchor], nums[explorer] = nums[explorer], nums[anchor]
			}
			anchor++
		}
	}
}

// @leet end

