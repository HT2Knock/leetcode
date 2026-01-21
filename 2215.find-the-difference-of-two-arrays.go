package leetcode

// @leet start
func findDifference(nums1 []int, nums2 []int) [][]int {
	s1, s2 := make(map[int]struct{}), make(map[int]struct{})

	for _, n := range nums1 {
		s1[n] = struct{}{}
	}
	for _, n := range nums2 {
		s2[n] = struct{}{}
	}

	result := make([][]int, 2)

	for k := range s1 {
		if _, ok := s2[k]; !ok {
			result[0] = append(result[0], k)
		}
	}

	for k := range s2 {
		if _, ok := s1[k]; !ok {
			result[1] = append(result[1], k)
		}
	}

	return result
}

// @leet end

