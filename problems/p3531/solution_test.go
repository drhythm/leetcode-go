package p3531

import "testing"


func TestCountCoveredBuildings(t *testing.T) {
	type args struct {
		n int
		buildings [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// Official Test
		{
			name: "Example 1",
			args: args{n: 3, buildings: [][]int{{1, 2},{2, 2},{3, 2},{2, 1},{2, 3}} },
			want: 1,
		},
		{
			name: "Example 2",
			args: args{n: 3, buildings: [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}},
			want: 0,
		},
		{
			name: "Example 3",
			args: args{n: 5, buildings: [][]int{{1, 3}, {3, 2}, {3, 3}, {3, 5}, {5, 3}}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVal := countCoveredBuildings(tt.args.n, tt.args.buildings); gotVal != tt.want {
				t.Errorf("countCoveredBuildings() %v, want %v", gotVal, tt.want)
				return
			}
		})
	}
}