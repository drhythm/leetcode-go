package p2211

import (
	"testing"
)

func Test_countCollisions(t *testing.T) {
	tests := [] struct {
		name	string
		args	string
		want	int
	}{
		// Official Tests
		{
			name:	"Example 1",
			args:	"RLRSLL",
			want:	5,
		},
		{
			name:	"Example 2",
			args:	"LLRR",
			want:	0,
		},
		// Corner Case
		{
			name:	"Internal Consecutive Cars Moving Right",
			args:	"SRRRRRS",
			want:	5,
		},
		{
			name:	"Internal Consecutive Cars Moving Left",
			args:	"SLLLLLS",
			want:	5,
		},
		{
			name:	"Internal Consecutive Cars Moving Left and right",
			args:	"SRRRLLS",
			want:	5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVal := countCollisions(tt.args); gotVal != tt.want {
				t.Errorf("countCollisions() = %v, want = %v\n", gotVal, tt.want)
				return
			}
		})
	}
}