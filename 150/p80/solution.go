package p80

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii

// removeDuplicates removes duplicates from a sorted slice in-place,
// keeping at most two occurrences of each value.
// It returns the new length (the first n elements of nums are the result).
func removeDuplicates(nums []int) int {
    // write is the index where the next kept element should be written.
    // Invariant: nums[:write] always satisfies the "at most two duplicates" rule.
    write := 0

    for _, x := range nums {
        // We can always keep the first two elements.
        // For the rest, we keep x only if it is greater than nums[write-2],
        // meaning we haven't already kept two copies of x.
        if write < 2 || x > nums[write-2] {
            nums[write] = x
            write++
        }
    }
    return write
}