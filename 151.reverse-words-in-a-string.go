package leetcode

// @leet start
func reverseWords(s string) string {
	b := []byte(s)
	n := len(b)

	reverse(b, 0, n-1)

	writeIdx := 0
	for i := 0; i < n; i++ {
		if b[i] != ' ' {
			if writeIdx != 0 {
				b[writeIdx] = ' '
				writeIdx++
			}

			start := writeIdx
			for i < n && b[i] != ' ' {
				b[writeIdx] = b[i]
				writeIdx++
				i++
			}

			reverse(b, start, writeIdx-1)
		}
	}

	return string(b[:writeIdx])
}

func reverse(b []byte, i, j int) {
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
}

// @leet end
