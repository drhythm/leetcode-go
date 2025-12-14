package p2147

import "testing"

func TestNumberOfWays(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		// Oficial Tests
		{
			name: "Example 1",
			args: "SSPPSPS",
			want: 3,
		},
		{
			name: "Example 2",
			args: "PPSPSP",
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVal := numberOfWays(tt.args); gotVal != tt.want {
				t.Errorf("numberOfWays() %v, want %v", gotVal, tt.want)
				return
			}
		})
	}
}
