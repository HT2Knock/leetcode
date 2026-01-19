package leetcode

import (
	"slices"
)

// @leet start
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	slices.Sort(nums)

	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		if nums[i] > 0 {
			break
		}

		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})

				for j < k && nums[j+1] == nums[j] {
					j++
				}
				for j < k && nums[k-1] == nums[k] {
					k--
				}
				j++
				k--
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}

	return result
}

// @leet end

