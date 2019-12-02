func lengthOfLongestSubstringTwoDistinct(s string) int {
    dict := make(map[byte]int)
    res := 0
    l := 0
    for r := 0; r < len(s); r++ {
        if _, ok := dict[s[r]]; ok {
            dict[s[r]] += 1
        } else {
            dict[s[r]] = 1
        }
        for l < r && len(dict) > 2 {
            dict[s[l]] -= 1
            if dict[s[l]] == 0 {
                delete(dict, s[l])
            }
            l += 1
        }
        res = max(res, r-l+1)
    }
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
