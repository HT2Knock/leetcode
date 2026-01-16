package leetcode

// @leet start
func compress(chars []byte) int {
	write := 0
	read := 0

	for read < len(chars) {
		currentChar := chars[read]
		groupLength := 0

		// 1. Count consecutive chars
		for read < len(chars) && chars[read] == currentChar {
			read++
			groupLength++
		}

		// 2. Write the current char
		chars[write] = currentChar
		write++

		// 3. Write the groupLength number
		if groupLength > 1 {
			anchor := write
			for groupLength > 0 {
				chars[write] = byte(groupLength%10 + '0')
				write++
				groupLength /= 10
			}
			inPlaceReverse(chars[anchor:write])
		}

	}

	return write
}

func inPlaceReverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// @leet end

