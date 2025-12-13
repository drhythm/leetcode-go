package p3606

import (
	"slices"
	"testing"
)

func TestValidateCoupons(t *testing.T) {
	type args struct {
		code         []string
		businessLine []string
		isActive     []bool
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		// Offical Tests
		{
			name: "Example 1",
			args: args{code: []string{"SAVE20","","PHARMA5","SAVE@20"},
				businessLine: []string{"restaurant","grocery","pharmacy","restaurant"},
				isActive:     []bool{true, true, true, true},
			},
			want: []string{"PHARMA5", "SAVE20"},
		},
		{
			name: "Example 2",
			args: args{code: []string{"GROCERY15","ELECTRONICS_50","DISCOUNT10"},
				businessLine: []string{"grocery","electronics","invalid"},
				isActive:     []bool{false, true, true},
			},
			want: []string{"ELECTRONICS_50"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVal := validateCoupons(tt.args.code, tt.args.businessLine, tt.args.isActive); !slices.Equal(gotVal, tt.want) {
				t.Errorf("validateCoupons() %v, want %v", gotVal, tt.want)
				return
			}
		})
	}
}
