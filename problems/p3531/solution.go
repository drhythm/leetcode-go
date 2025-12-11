package p3531

import "math"

// https://leetcode.cn/problems/count-covered-buildings
func countCoveredBuildings(n int, buildings [][]int)  int {
    rows := make([]Bounds, n+1)
    cols := make([]Bounds, n+1)

    for i := 0; i <= n; i++ {
        rows[i] = Bounds{Min: math.MaxInt, Max: math.MinInt}
        cols[i] = Bounds{Min: math.MaxInt, Max: math.MinInt}
    }

    for _, p := range buildings {
        x, y := p[0], p[1]
        rows[x].Update(y)
        cols[y].Update(x)
    }
    count := 0
    for _, p := range buildings {
        x, y := p[0], p[1]
        if rows[x].Covers(y) && cols[y].Covers(x) {
            count++
        }
    }
    return count
}


type Bounds struct {
    Max int
    Min int
}

func (b *Bounds) Update(x int) {
	if b.Max < x {
		b.Max = x
	}
	if b.Min > x {
		b.Min = x
	}
}

func (b *Bounds) Covers(val int) bool {
    return b.Min < val && val < b.Max
}