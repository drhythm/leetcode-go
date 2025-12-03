package p3625

import (
	"testing"
)

func Test_countTrapezoids(t *testing.T) {
	tests := []struct{
		name	string
		args	[][]int
		want	int
	}{
		// Official Tests
		{
			name:	"Example 1",
			args:	[][]int{{-3, 2}, {3, 0}, {2, 3}, {3, 2}, {2, -3}},
			want: 	2,
		},
		{
			name:	"Example 2",
			args:	[][]int{{0, 0}, {1, 0}, {0, 1}, {2, 1}},
			want:	1,
		},
		// Conner Case:
		{
			name:	"Parallelogram",
			args:	[][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
			want:	1,
		},
		{
			name:	"Collinear Segment and Sharing Midpoint",
			args:	[][]int{{0, -1}, {0, 1}, {0, 2}, {0, -2}},
			want:	0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			if gotVal := countTrapezoids(tt.args); gotVal != tt.want {
				t.Errorf("countTrapezoids() = %v, tt.want %v\n", gotVal, tt.want)
				return
			}
		})
	}
}