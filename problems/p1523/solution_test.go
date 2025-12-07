package p1523

import (
	"testing"
)

func Test_countOdds(t *testing.T) {
	type args struct {
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// Official Tests
		{
			name: "Example 1",
			args: args{low: 3, high: 7},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{low: 8, high: 10},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVal := countOdds(tt.args.low, tt.args.high); gotVal != tt.want {
				t.Errorf("countOdds() = %v, tt.want = %v", gotVal, tt.want)
				return
			}
		})
	}
}
