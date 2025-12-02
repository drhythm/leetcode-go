package p3623

//https://leetcode.cn/problems/count-number-of-trapezoids-i
func countTrapezoids(points [][]int) int {
    const mod int = 1_000_000_007
    // Count points for each y-coordinate
    // yCounts -> key: y-coordinate, value: number of points at this y-level
    yCounts := make(map[int]int)
    for _, p := range points {
        // p[1] represents the y-coordinate
        yCounts[p[1]] += 1
    }
    var totalTrapezoids, prefixPairs int
    // iterate over the counts of points at each y-level
    for _, count := range yCounts {
        if count < 2 {
            continue
        }
        // Calculate the number of pairs on the current line: C(n, 2)
        // Formula: n * (n - 1) / 2
        pairs := count * (count - 1) /2
        totalTrapezoids = (totalTrapezoids + (prefixPairs * pairs)) % mod
        prefixPairs = (prefixPairs + pairs) % mod
    }
    return totalTrapezoids
}