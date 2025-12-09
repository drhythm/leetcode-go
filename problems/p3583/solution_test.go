package p3583

import (
	"testing"
)

func TestSpecialTriplets(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		// Official Tests
		{
			name: "Example 1",
			args: []int{6, 3, 6},
			want: 1,
		},
		{
			name: "Example 2",
			args: []int{0, 1, 0, 0},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVal := specialTriplets(tt.args); gotVal != tt.want {
				t.Errorf("specialTriplets() %v, want %v", gotVal, tt.want)
				return
			}
		})
	}
}
