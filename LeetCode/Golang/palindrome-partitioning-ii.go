func minCut(s string) int {
    cut := make([]int, len(s))
    dp := make([][]bool, len(s))
    for i := 0; i < len(s); i++ {
        dp[i] = make([]bool, len(s))
    }
    for i := 0; i < len(s); i++ {
        cut[i] = i
        for j := 0; j <= i; j++ {
            if (s[i] == s[j] && (j+1 > i-1 || dp[j+1][i-1])) {
                dp[j][i] = true
                if (j == 0) {
                    cut[i] = 0
                } else if (cut[j-1]+1 < cut[i]) {
                    cut[i] = cut[j-1] + 1
                }
            }
        }
    }
    fmt.Println(dp)
    return cut[len(s)-1]
}
