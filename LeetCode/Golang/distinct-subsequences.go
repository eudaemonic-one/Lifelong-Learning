func numDistinct(s string, t string) int {
    dp := make([][]int, len(s)+1)
    for i := 0; i < len(s)+1; i++ {
        dp[i] = make([]int, len(t)+1)
    }
    for i := range s {
        dp[i][0] = 1
    }
    for i, c := range s {
        for j, v := range t {
            if c == v {
                dp[i+1][j+1] = dp[i][j] + dp[i][j+1]
            } else {
                dp[i+1][j+1] = dp[i][j+1]
            }
        }
    }
    return dp[len(s)][len(t)]
}