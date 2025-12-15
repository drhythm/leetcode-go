package p2110

// https://leetcode.cn/problems/number-of-smooth-descent-periods-of-a-stock
func getDescentPeriods(prices []int) int64 {
    length := 1
    ans := 1
    for i := 1; i < len(prices); i++ {
        if prices[i] == prices[i-1]-1 {
            length++
        } else {
            length = 1
        }
        ans += length
    }
    return int64(ans)
}