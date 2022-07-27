package sol

func checkValidString(s string) bool {
	sLen := len(s)
	maxLeft, minLeft := 0, 0
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for pos := 0; pos < sLen; pos++ {
		ch := s[pos]
		if ch != ')' {
			maxLeft += 1
		} else {
			maxLeft -= 1
		}
		if ch != '(' {
			minLeft -= 1
		} else {
			minLeft += 1
		}
		minLeft = max(0, minLeft)
		if maxLeft < 0 {
			return false
		}
	}
	return minLeft == 0
}
