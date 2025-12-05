package p3432

// https://leetcode.cn/problems/count-partitions-with-even-sum-difference
func countPartitions(nums []int) int {
    // Edge Case Check
    if len(nums) < 2 {
        return 0
    }
    // Calculate Total Sum
    sum := 0
    for _, num := range nums {
        sum += num
    }
    // Case 1: Odd
    // Odd = Odd + Even
    if sum % 2 != 0 {
        return 0
    }
    // Case 2: Even
    // Even = Even + Even(or Odd + Odd)
    return len(nums) - 1
}