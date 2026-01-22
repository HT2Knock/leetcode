package leetcode

// @leet start
func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	for _, curr := range asteroids {
		isDestroyed := false

		for len(stack) > 0 && curr < 0 && stack[len(stack)-1] > 0 {
			top := stack[len(stack)-1]

			if top < -curr {
				stack = stack[:len(stack)-1]
				continue
			} else if top == -curr {
				stack = stack[:len(stack)-1]
			}

			isDestroyed = true
			break
		}

		if !isDestroyed {
			stack = append(stack, curr)
		}
	}

	return stack
}

// @leet end

