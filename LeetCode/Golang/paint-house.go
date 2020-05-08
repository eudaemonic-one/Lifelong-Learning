func minCost(costs [][]int) int {
    if costs == nil || len(costs) == 0 {
        return 0
    }
    for i := 1; i < len(costs); i++ {
        costs[i][0] += min(costs[i-1][1],costs[i-1][2])
        costs[i][1] += min(costs[i-1][0],costs[i-1][2])
        costs[i][2] += min(costs[i-1][1],costs[i-1][0])
    }
    return min(costs[len(costs)-1][0], min(costs[len(costs)-1][1], costs[len(costs)-1][2]))
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
