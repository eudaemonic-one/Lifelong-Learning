func maxProfit(prices []int) int {
    if len(prices) <= 1 {
        return 0
    }
    N := len(prices)
    have1sell0 := -prices[0]
    have1sell1 := 0
    have0buy0 := 0
    have0buy1 := -prices[0]
    for i := 1; i < N; i++ {
        have1sell0 = max(have1sell0, have0buy1)
        have0buy1 = -prices[i] + have0buy0
        have0buy0 = max(have0buy0, have1sell1)
        have1sell1 = prices[i] + have1sell0
    }
    return max(have1sell1, have0buy0)
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
