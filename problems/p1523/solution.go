package p1523

// https://leetcode.cn/problems/count-odd-numbers-in-an-interval-range
func countOdds(low int, high int) int {
	// head and tail are all even. (head)even .... even(tail)
	if low&1 == 0 && high&1 == 0 {
		return (high - low) / 2
	}
	// else
	return (high-low)/2 + 1
}
