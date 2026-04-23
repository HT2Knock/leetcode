package leetcode

// @leet start
func isValid(s string) bool {
	stack := make([]byte, 0, len(s)/2)
	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}

	for i := 0; i < len(s); i++ {
		char := s[i]
		if expected, exist := pairs[char]; exist {
			if len(stack) == 0 || stack[len(stack)-1] != expected {
				return false
			}

			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, char)
	}

	return len(stack) == 0
}

// @leet end

