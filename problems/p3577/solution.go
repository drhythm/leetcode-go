package p3577

// https://leetcode.cn/problems/count-the-number-of-computer-unlocking-permutations
func countPermutations(complexity []int) (ans int) {
	const mod = 1_000_000_007
	n := len(complexity)
	ans = 1
	for i := 1; i < n; i++ {
		if complexity[i] <= complexity[0] {
			return 0
		}
		ans = ans * i % mod
	}
	return ans
}
