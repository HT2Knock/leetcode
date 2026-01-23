package leetcode

// @leet start
func predictPartyVictory(senate string) string {
	queue := []byte(senate)

	rCount, dCount := 0, 0
	for i := 0; i < len(senate); i++ {
		if senate[i] == 'R' {
			rCount++
		} else {
			dCount++
		}
	}

	floatingRadiantBans, floatingDireBans := 0, 0
	for i := 0; rCount > 0 && dCount > 0; i++ {
		if queue[i] == 'R' {
			if floatingRadiantBans > 0 {
				floatingRadiantBans--
				rCount--
				continue
			}

			floatingDireBans++
			queue = append(queue, queue[i])
		} else {
			if floatingDireBans > 0 {
				floatingDireBans--
				dCount--
				continue
			}

			floatingRadiantBans++
			queue = append(queue, queue[i])
		}
	}

	if rCount > 0 {
		return "Radiant"
	}

	return "Dire"
}

// @leet end

