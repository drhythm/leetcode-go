package equalmap

import (
	"math"
	"testing"
)

func TestEqual(t *testing.T) {
	tests := []struct {
		name string
		a    map[string]int
		b    map[string]int
		want bool
	}{
		{
			name: "equal maps",
			a:    map[string]int{"x": 1, "y": 2},
			b:    map[string]int{"x": 1, "y": 2},
			want: true,
		},
		{
			name: "different values",
			a:    map[string]int{"x": 1, "y": 2},
			b:    map[string]int{"x": 1, "y": 3},
			want: false,
		},
		{
			name: "different keys",
			a:    map[string]int{"x": 1, "y": 2},
			b:    map[string]int{"x": 1, "z": 2},
			want: false,
		},
		{
			name: "different lengths",
			a:    map[string]int{"x": 1, "y": 2},
			b:    map[string]int{"x": 1},
			want: false,
		},
		{
			name: "both empty",
			a:    map[string]int{},
			b:    map[string]int{},
			want: true,
		},
		{
			name: "one empty",
			a:    map[string]int{"x": 1},
			b:    map[string]int{},
			want: false,
		},
		{
			name: "nil and empty",
			a:    map[string]int{},
			b:    nil,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.a, tt.b); got != tt.want {
				t.Errorf("EqualMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqual_NaN(t *testing.T) {
	a := map[string]float64{"val": math.NaN()}
	b := map[string]float64{"val": math.NaN()}
	if Equal(a, b) {
		t.Errorf("Expected false for NaN comparison using ==, but got true")
	}
}