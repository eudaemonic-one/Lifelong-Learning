func shortestDistance(words []string, word1 string, word2 string) int {
    idx1, idx2 := -1, -1
    res := len(words) + 1
    for i, word := range words {
        if word == word1 {
            idx1 = i
            if idx1 != -1 && idx2 != -1 {
                res = min(res, diff(idx1, idx2))
            }
        } else if word == word2 {
            idx2 = i
            if idx1 != -1 && idx2 != -1 {
                res = min(res, diff(idx1, idx2))
            }
        }
    }
    return res
}

func diff(x, y int) int {
    diff := x - y
    if diff < 0 {
        return -diff
    }
    return diff
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
