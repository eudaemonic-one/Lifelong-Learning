func minCostII(costs [][]int) int {
    N := len(costs)
    if N == 0 {
        return 0
    }
    K := len(costs[0])
    if K == 0 {
        return 0
    }
    const INT_MAX int = int(^uint(0) >> 1)
    for i := 1; i < N; i++ {
        idx1, min1 := 0, INT_MAX
        _, min2 := 0, INT_MAX
        for k := 0; k < K; k++ {
            if costs[i-1][k] < min1 {
                _, min2 = idx1, min1
                idx1, min1 = k, costs[i-1][k]
            } else if costs[i-1][k] < min2 {
                _, min2 = k, costs[i-1][k]
            }
        }
        for k := 0; k < K; k++ {
            if k == idx1 {
                costs[i][k] += min2
            } else {
                costs[i][k] += min1
            }
        }
    }
    minCost := INT_MAX
    for k := 0; k < K; k++ {
        minCost = min(minCost, costs[N-1][k])
    }
    return minCost
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
