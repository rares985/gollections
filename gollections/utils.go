package gollections

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max3(a, b, c int) int {
	return max(max(a, b), c)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	return min(min(a, b), c)
}
