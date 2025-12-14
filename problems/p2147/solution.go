package p2147

// https://leetcode.cn/problems/number-of-ways-to-divide-a-long-corridor

func numberOfWays(corridor string) int {
	const mod = 1_000_000_007
	ans, cntS, lastS := 1, 0, 0
	for i, ch := range corridor {
		if ch == 'S' {
			cntS++
			if cntS >= 3 && cntS%2 > 0 {
				ans = ans * (i - lastS) % mod
			}
			lastS = i
		}
	}
	if cntS == 0 || cntS%2 > 0 {
		return 0
	}
	return ans
}
