package p2141


// https://leetcode.com/problems/maximum-running-time-of-n-computers
func maxRunTime(n int, batteries []int) int64 {
    // 1. Calculate total battery power using int64 to prevent overflow.
    var totalPower int64
    for _, b := range batteries {
        totalPower += int64(b)
    }

    // 2. Binary Search Range
    // left: min possible time (0)
    // right: theoretical max time (total power / computers)
    var left int64 = 0
    var right int64 = totalPower / int64(n)
    
    // Convert n to int64 to avoid repeated casting inside the loop
    n64 := int64(n)

    // 3. The Check Function (Closure)
    // Returns true if 'n' computers can run for 'targetTime'
    check := func(targetTime int64) bool {
        var currentSum int64
        for _, b := range batteries {
            val := int64(b)
            // Logic: A battery can contribute at most 'targetTime'
            if val < targetTime {
                currentSum += val
            } else {
                currentSum += targetTime
            }
        }
        // Check if effective power is enough to cover the demand (n * targetTime)
        return currentSum >= n64 * targetTime
    }

    // 4. Binary Search (Finding the Maximum satisfying value)
    for left < right {
        // Bias mid to the right to prevent infinite loop when left = right - 1
        mid := (right-left+1)/2 + left
        
        if check(mid) {
            // mid is possible, so the answer is at least mid.
            // Try to find a larger one in [mid, right]
            left = mid
        } else {
            // mid is impossible, the answer must be in [left, mid-1]
            right = mid - 1
        }
    }

    return left
}