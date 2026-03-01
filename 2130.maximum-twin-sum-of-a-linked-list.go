package leetcode

// @leet start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode
	cur := slow
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}

	first, second := head, prev
	maxSum := 0
	for second != nil {
		sum := first.Val + second.Val
		if sum > maxSum {
			maxSum = sum
		}

		first = first.Next
		second = second.Next
	}

	return maxSum
}

// @leet end
