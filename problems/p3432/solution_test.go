package p3432

import (
	"testing"
)

func Test_countPartitions(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		// Official Tests
		{
			name: "Example 1",
			args: []int{10, 10, 3, 7, 6},
			want: 4,
		},
		{
			name: "Example 2",
			args: []int{1, 2, 2},
			want: 0,
		},
		{
			name: "Example 3",
			args: []int{2, 4, 6, 8},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVal := countPartitions(tt.args); gotVal != tt.want {
				t.Errorf("countPartitions() = %v, tt.want = %v\n", gotVal, tt.want)
				return
			}
		})
	}
}
