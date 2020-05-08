func wordBreak(s string, wordDict []string) []string {
    res := make([]string, 0)
    if len(s) == 0 {
        return res
    }
    dict := make(map[string]int)
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
    if !dp[len(s)-1] {
        return res
    }
    backtrack(&res, s, dict, []string{}, 0)
    return res
}

func backtrack(res *[]string, s string, dict map[string]int, path []string, i int) {
    if i == len(s) {
        *res = append(*res, strings.Join(path, " "))
        return
    }
    for j := i ; j < len(s); j++ {
        if _, ok := dict[s[i:j+1]]; ok {
            path = append(path, s[i:j+1])
            backtrack(res, s, dict, path, j+1)
            path = path[:len(path)-1]
        }
    }
}
