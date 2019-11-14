func minimumTotal(triangle [][]int) int {
    dp := make([]int, len(triangle))
    for i := range triangle {
        dp[i] = triangle[len(triangle)-1][i]
    }
    for i := len(triangle)-2; i >= 0; i-- {
        for j := 0; j < len(triangle[i]); j++ {
            if dp[j] < dp[j+1] {
                dp[j] = dp[j] + triangle[i][j]
            } else {
                dp[j] = dp[j+1] + triangle[i][j]
            }
        }
    }
    return dp[0]
}