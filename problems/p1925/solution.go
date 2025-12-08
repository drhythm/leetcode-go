package p1925

import "math"

// https://leetcode.cn/problems/count-square-sum-triples
func countTriples(n int) (ans int) {
	for a := 1; a < n; a++ {
		for b := 1; b < a && a*a+b*b <= n*n; b++ {
			c2 := a*a + b*b
			c := int(math.Sqrt(float64(c2)))
			if c*c == c2 {
				ans++
			}
		}
	}
	return ans * 2
}
