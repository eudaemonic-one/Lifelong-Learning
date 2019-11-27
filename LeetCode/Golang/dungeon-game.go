func calculateMinimumHP(dungeon [][]int) int {
    const INT_MAX int = int(^uint(0) >> 1)
    m := len(dungeon)
    if m == 0 {
        return 0
    }
    n := len(dungeon[0])
    if n == 0 {
        return 0
    }
    dp := make([][]int, m+1)
    for i := 0; i < m+1; i++ {
        dp[i] = make([]int, n+1)
        for j := 0; j < n+1; j++ {
            dp[i][j] = INT_MAX
        }
    }
    dp[m][n-1] = 1
    dp[m-1][n] = 1
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            need := min(dp[i+1][j], dp[i][j+1]) - dungeon[i][j]
            if need <= 0 {
                dp[i][j] = 1
            } else {
                dp[i][j] = need
            }
        }
    }
    return dp[0][0]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
