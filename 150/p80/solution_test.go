package p80

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		args []int
		expected int
		want []int
	}{
		{
			name: "given_case_1",
			args: []int{1,1,1,2,2,3},
			expected: 5,
			want: []int{1,1,2,2,3},
		},
		{
			name: "given_case_2",
			args: []int{0,0,1,1,1,1,2,3,3},
			expected: 7,
			want: []int{0,0,1,1,2,3,3},
		},
		{
			name: "empty",
			args: []int{},
			expected: 0,
			want: []int{},
		},
		{
			name: "single_element",
			args: []int{1},
			expected: 1,
			want: []int{1},
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.args)
			if got != tt.expected {
				t.Fatalf("removeDuplicates() = %d, want %d", got, tt.expected)
			}
			for i, v := range tt.want {
				if tt.args[i] != v {
					t.Fatalf("nums[%d] = %d, want %d; nums=%v, want=%v",
						i, tt.args[i], v, tt.args[:got], tt.want)
				}
			}
		})
	}
}