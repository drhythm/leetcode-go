package p2141

import (
	"testing"
)

func Test_maxRunTime(t *testing.T) {
	type args struct {
		n int
		batteries []int
	}
	tests := []struct {
		name	string
		args	args
		want	int64
	}{
		// Official Tests
		{
			name:	"Example 1",
			args:	args{n: 2, batteries: []int{3, 3, 3}},
			want:	4,
		},
		{
			name:	"Example 2",
			args:	args{n: 2, batteries: []int{1, 1, 1, 1}},
			want:	2,
		},
		// Edge Case: N = 1(one computer)
		// With only one computer, the result is the sum of all batteries.
		{
			name:	"One Computer",
			args:	args{n: 1, batteries: []int{2, 5, 6, 9}},
			want:	22,
		},
		// Overflow Case: Integer Overflow
		// Ensures int64 is correctly for large sums.
		{
			name:	"Integer Overflow",
			args:	args{n: 2 , batteries: []int{10_000_000_000, 10_000_000_000, 10_000_000_000, 10_000_000_000}},
			want:	20_000_000_000,
		},
		// Corner Case: Dominant Batteries
		// A huge battery can only be used for single computer.
		{
			name:	"Corner Case",
			args:	args{n: 2, batteries: []int{500, 1, 1}},
			want:	2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			if got := maxRunTime(tt.args.n, tt.args.batteries); got != tt.want {
				t.Errorf("maxRunTime() = %v, want = %v", got, tt.want)
			}
		})
	}
}