package p27

import (
	"testing"

	"github.com/drhythm/leetcode-go/common/equalmap"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		expected int
		want     map[int]int
	}{
		{
			name:     "remove from middle",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			expected: 2,
			want:     map[int]int{2: 2},
		},
		{
			name:     "remove multiple elements",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			expected: 5,
			want:     map[int]int{0: 2, 1: 1, 3: 1, 4: 1},
		},
		{
			name:     "no elements to remove",
			nums:     []int{1, 2, 3},
			val:      4,
			expected: 3,
			want:     map[int]int{1: 1, 2: 1, 3: 1},
		},
		{
			name:     "remove all elements",
			nums:     []int{5, 5, 5},
			val:      5,
			expected: 0,
			want:     map[int]int{},
		},
		{
			name:     "empty slice",
			nums:     []int{},
			val:      1,
			expected: 0,
			want:     map[int]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.nums, tt.val)
			if got != tt.expected {
				t.Errorf("removeElement() = %d, want %d", got, tt.expected)
			}
			cnt := map[int]int{}
			for i := 0; i < tt.expected; i++ {
				cnt[tt.nums[i]]++
			}
			if equalmap.Equal(cnt, tt.want) {
				t.Errorf("removeElement() = %v, want %v", cnt, tt.want)
			}
		})
	}
}
