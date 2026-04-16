package leetcode

// @leet start
func totalFruit(fruits []int) int {
	if len(fruits) < 2 {
		return len(fruits)
	}

	start := 0
	maxLength := 0
	basket := make(map[int]int, 3)

	for end, fruit := range fruits {
		basket[fruit]++

		for len(basket) > 2 {
			leftFruit := fruits[start]

			basket[leftFruit]--
			if basket[leftFruit] == 0 {
				delete(basket, leftFruit)
			}
			start++
		}

		currentLength := end - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

// @leet end
