package leetcode

import (
	"slices"
)

// @leet start
func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			switch {
			case sum < 0:
				left++

			case sum > 0:
				right--

			default:
				res = append(res, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			}
		}
	}

	return res
}

// @leet end
