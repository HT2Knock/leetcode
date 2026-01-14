package leetcode

// @leet start
import "slices"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	maxVal := slices.Max(candies)

	threshold := maxVal - extraCandies

	for i, candyCount := range candies {
		result[i] = candyCount >= threshold
	}

	return result
}

// @leet end
