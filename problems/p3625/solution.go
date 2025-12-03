package p3625

// https://leetcode.cn/problems/count-number-of-trapezoids-ii

// GCD helper to handle slope normalization
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Slope uses reduced fraction dy/dx to represent slope precisely.
type Slope struct {
	dy int
	dx int
}

// Point represents a coordinate.
type Point struct {
	x int
	y int
}

func countTrapezoids(points [][]int) (ans int) {
	n := len(points)
	if n < 4 {
		return 0
	}

	// Group 1: Identify parallel segments.
	// Map: Slope -> [List of Line Equations(Intercepts)]
	// Key: Slope, Value: slice of identifiers for the specific line equation
	parallelGroups := make(map[Slope][]int)

	// Group 2: Identify parallelograms via diagonals.
	// Map: Midpoint -> (Map: Slope -> Count)
	// We need to group by Slope INSIDE the midpoint map to avoid subtracting collinear lines.
	midpointSlopeCounts := make(map[Point]map[Slope]int)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			x1, y1 := points[i][0], points[i][1]
			x2, y2 := points[j][0], points[j][1]

			// 1. Calculate Slope (dy/dx)
			dx := x1 - x2
			dy := y1 - y2
			
			// Normalize slope using GCD to avoid float precision issues
			if dx == 0 {
				dy = 1
			} else if dy == 0 {
				dx = 1
			} else {
				if dx < 0 {
					dx = -dx
					dy = -dy
				}
				g := gcd(abs(dx), abs(dy))
				dx /= g
				dy /= g
			}
			slope := Slope{dy, dx}

			// 2. Add to Parallel Groups
			// 'val' represents the line equation constant: C = -A*x - B*y
			// val distinguishs parallel lines from collinear lines.
			val := -dy*x1 + dx*y1
			parallelGroups[slope] = append(parallelGroups[slope], val)

			// 3. Add to Midpoint Groups
			// Midpoint * 2 = (x1+x2, y1+y2) to keep integers
			mid := Point{x1 + x2, y1 + y2}
			
			// Initialize inner map if not exists
			if midpointSlopeCounts[mid] == nil {
				midpointSlopeCounts[mid] = make(map[Slope]int)
			}
			midpointSlopeCounts[mid][slope]++
		}
	}

	// 4. Calculate total parallel pairs (Potential Trapezoids)
	for _, intercepts := range parallelGroups {
		if len(intercepts) < 2 {
			continue
		}
		
		// Count lines that are collinear (share the same intercept 'val')
		// We only want to pair lines with DIFFERENT 'val'.
		lineCounts := make(map[int]int)
		for _, val := range intercepts {
			lineCounts[val]++
		}

		// Calculate pairs: Previous Sum * Current Count
		totalSegments := 0
		for _, count := range lineCounts {
			ans += totalSegments * count
			totalSegments += count
		}
	}

	// 5. Subtract Parallelograms
	// A parallelogram is formed by diagonals bisecting each other with DIFFERENT slopes.
	for _, slopesMap := range midpointSlopeCounts {
		if len(slopesMap) < 2 {
			continue
		}
		
		// Logic: Only pair segments with DIFFERENT slopes at the same midpoint.
		// Collinear segments (Same Slope, Same Midpoint) are not parallelograms.
		totalDiagonals := 0
		for _, count := range slopesMap {
			ans -= totalDiagonals * count
			totalDiagonals += count
		}
	}

	return ans
}