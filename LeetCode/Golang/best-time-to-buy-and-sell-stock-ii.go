func maxProfit(prices []int) int {
    var min, ans = int(^uint(0) >> 1), 0
    for _, price := range(prices) {
        if (price < min) {
            min = price
        } else {
            ans += price - min
            min = price
        }
    }
    return ans
}