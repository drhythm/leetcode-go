package p3578

import (
	"testing"
)

func Test_countPartitions(t *testing.T) {
	const mod int = 1_000_000_007
	type args struct {
		vals []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// Official Tests
		{
			name: "Example 1",
			args: args{vals: []int{9, 4, 1, 3, 7}, k: 4},
			want: 6,
		},
		{
			name: "Example 2",
			args: args{vals: []int{3, 3, 4}, k: 0},
			want: 2,
		},
		// Zig-zag Case
		{
			name: "Zig-zag Data",
			args: args{vals: []int{1, 10, 1, 10}, k: 5},
			want: 1,
		},
		// Modulo Arithmetic Test
		{
			name: "Modulo Check",
			args: args{
				vals: func() []int {
					arr := make([]int, 100)
					for index, _ := range arr {
						arr[index] = 6
					}
					return arr
				}(),
				k: 0,
			},
			want: func() int {
				ans := 1
				for i := 0; i < 99; i++ {
					ans = (ans * 2) % mod
				}
				return ans
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVal := countPartitions(tt.args.vals, tt.args.k); gotVal != tt.want {
				t.Errorf("countPartitions() = %v, tt.want = %v", gotVal, tt.want)
				return
			}
		})
	}
}
