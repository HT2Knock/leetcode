package leetcode

import (
	"strings"
)

// @leet start
func decodeString(s string) string {
	numStack := []int{}
	strStack := []string{}
	currentNum := 0
	currentStr := ""

	for _, char := range s {
		if char >= '0' && char <= '9' {
			currentNum = currentNum*10 + int(char-'0')
		} else if char == '[' {
			numStack = append(numStack, currentNum)
			strStack = append(strStack, currentStr)

			currentStr = ""
			currentNum = 0
		} else if char == ']' {
			count := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			prevStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			currentStr = prevStr + strings.Repeat(currentStr, count)
		} else {
			currentStr += string(char)
		}
	}

	return currentStr
}

// @leet end
