package p3578

// https://leetcode.cn/problems/count-partitions-with-max-min-difference-at-most-k
func countPartitions(nums []int, k int) int {
	const mod = 1_000_000_007
	n := len(nums)

	// minQ and maxQ are monotonic queues storing indices.
	// minQ maintains indices with values in increasing order.
	// maxQ maintains indices with values in decreasing order.
	var minQ, maxQ []int

	// f[i] stores the number of valid partitions for the prefix of length i.
	// Base case: f[0] = 1 represents an empty partition.
	f := make([]int, n+1)
	f[0] = 1

	// prefixSum maintains the sum of valid f[j] within the current sliding window [left, i].
	prefixSum := 0
	left := 0

	for i, x := range nums {
		// Accumulate the current DP state into the running sum.
		prefixSum = (prefixSum + f[i]) % mod

		// Pop elements from back that are greater than or equal to the current element.
		for len(minQ) > 0 && x <= nums[minQ[len(minQ)-1]] {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, i)

		// Pop elements from back that are smaller than or equal to the current element.
		for len(maxQ) > 0 && x >= nums[maxQ[len(maxQ)-1]] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, i)

		// Shrink the sliding window from the left.
		// If the difference between max and min in the current window exceeds k,
		// we must move the left pointer forward to exclude invalid partitions.
		for nums[maxQ[0]]-nums[minQ[0]] > k {
			// Remove the invalid DP state from running sum.
			prefixSum = (prefixSum - f[left] + mod) % mod

			left++

			// Remove indices out of the current window (index < left).
			if minQ[0] < left {
				minQ = minQ[1:]
			}
			if maxQ[0] < left {
				maxQ = maxQ[1:]
			}
		}

		// Indices from left to i are all valid partitions.
		// f[i+1] = f[left] + ... + f[i] (which is prefixSum)
		f[i+1] = prefixSum
	}

	return f[n]
}
