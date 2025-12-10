package p3577

import "testing"

func TestCountPermutations(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		// Official Tests
		{
			name: "Example 1",
			args: []int{1, 2, 3},
			want: 2,
		},
		{
			name: "Example 2",
			args: []int{3, 3, 3, 4, 4, 4},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVal := countPermutations(tt.args); gotVal != tt.want {
				t.Errorf("countPermutations() %v, want %v", gotVal, tt.want)
				return
			}
		})
	}
}