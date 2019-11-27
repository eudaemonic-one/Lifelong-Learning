func wordBreak(s string, wordDict []string) bool {
    if len(s) == 0 {
        return false
    }
    dict := make(map[string]int, 0)
    for i := range wordDict {
        dict[wordDict[i]] = 1
    }
    dp := make([]bool, len(s))
    for j := 0; j < len(s); j++ {
        for i := 0; i <= j; i++ {
            if _, ok := dict[s[i:j+1]]; (i == 0 || dp[i-1]) && ok {
                dp[j] = true
            }
        }
    }
    return dp[len(s)-1]
}
