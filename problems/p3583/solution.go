package p3583

// https://leetcode.cn/problems/count-special-triplets/description
func specialTriplets(nums []int) (ans int) {
	const mod = 1_000_000_007
	// cnti stores frequency of each number encountered so far.
	cnti := make(map[int]int)
	// cntij stores the frequency of valid pairs (i, j).
	cntij := make(map[int]int)
	for _, num := range nums {
		if num%2 == 0 {
			// accumulate the count of matched triplets.
			ans += cntij[num/2]
		}
		cntij[num] += cnti[num*2]
		cnti[num]++
	}
	return ans % mod
}
