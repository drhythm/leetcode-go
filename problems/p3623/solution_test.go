package p3623

import (
	"testing"
)

func Test_countTrapezoids(t *testing.T) {
	generateLargeData := func(n int) [][]int {
		points := make([][]int, 0, 2*n)
		for i := 0; i < n; i++ {
			points = append(points, []int{0, 1})
		}
		for i := 0; i < n; i++ {
			points = append(points, []int{0, 2})
		}
		return points
	}
	calculateExpected := func(n int) int {
		const mod = 1_000_000_007
		pairs := int64(n) * int64(n-1) / 2
		result := (pairs % mod) * (pairs % mod)
		return int(result % mod)
	}
	// 40000
	// 40000 * 39999 / 2 ≈ 800,000,000 (Close to 10^9)
	// 800M * 800M ≈ 6.4 * 10^17 (Fits in int64, tests overflow logic)
	LargeN := 40_000
	tests := []struct {
		name string
		args [][]int
		want int
	}{
		// Official Tests
		{
			name: "Example 1",
			args: [][]int{{1, 0}, {2, 0}, {3, 0}, {2, 2}, {3, 2}},
			want: 3,
		},
		{
			name: "Example 2",
			args: [][]int{{0, 0}, {1, 0}, {0, 1}, {2, 1}},
			want: 1,
		},
		// Overflow Case: Integer Overflow
		{
			name: "Integer Overflow",
			args: generateLargeData(LargeN),
			want: calculateExpected(LargeN),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVal := countTrapezoids(tt.args); gotVal != tt.want {
				t.Errorf("countTrapezoids() %v, tt.want %v\n", gotVal, tt.want)
				return
			}
		})
	}
}
