package p1925

import "testing"

func TestCountTriples(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		// Official Tests
		{
			name: "Example 1",
			args: 5,
			want: 2,
		},
		{
			name: "Example 2",
			args: 10,
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVal := countTriples(tt.args); gotVal != tt.want {
				t.Errorf("countTriples() %v, want %v", gotVal, tt.want)
				return
			}
		})
	}
}
