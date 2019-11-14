func maxProfit(prices []int) int {
    if (len(prices) == 0) {
        return 0
    }
    var min1, min2 = prices[0], prices[0]
    var profit1, profit2 = 0, 0
    for i := 1; i < len(prices); i++ {
        if (prices[i] < min1) {
            min1 = prices[i]
        }
        if (prices[i] - min1 > profit1) {
            profit1 = prices[i] - min1
        }
        if (prices[i] - profit1 < min2) {
            min2 = prices[i] - profit1
        }
        if (prices[i] - min2 > profit2) {
            profit2 = prices[i] - min2
        }
    }
    return profit2
}