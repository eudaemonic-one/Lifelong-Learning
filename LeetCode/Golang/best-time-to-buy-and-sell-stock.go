func maxProfit(prices []int) int {
    const INT_MAX = int(^uint(0) >> 1)
    var minPrice = INT_MAX
    var ans = 0
    for i := 0; i < len(prices); i++ {
        if (prices[i] < minPrice) {
            minPrice = prices[i]
        } else if (prices[i] - minPrice > ans) {
            ans = prices[i] - minPrice
        }
    }
    return ans
}