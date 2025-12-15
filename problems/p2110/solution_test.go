package p2110

import "testing"

func TestGetDescentPeriods(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int64
	}{
		// Oficial Testing
		{
			name: "Example 1",
			args: []int{3,2,1,4},
			want: 7,
		},
		{
			name: "Example 2",
			args: []int{8,6,7,7},
			want: 4,
		},
		{
			name: "Example 3",
			args: []int{1},
			want: 1,
		},
	}
	for _, tt := range tests {
		if gotVal := getDescentPeriods(tt.args); gotVal != tt.want {
			t.Errorf("getDescentPeriods() %v, want %v", gotVal, tt.want)
			return
		}
	}
}