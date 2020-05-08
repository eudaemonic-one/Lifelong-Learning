func shortestWordDistance(words []string, word1 string, word2 string) int {
    distMap := make(map[string][]int)
    for i, word := range words {
        distMap[word] = append(distMap[word], i)
    }
    loc1, loc2 := distMap[word1], distMap[word2]
    l1, l2 := 0, 0
    res := int(^uint(0) >> 1)
    if word1 == word2 {
        for i := 0; i < len(loc1)-1; i++ {
            res = min(res, loc1[i+1]-loc1[i])
        }
        return res
    }
    for l1 < len(loc1) && l2 < len(loc2) {
        res = min(res, abs(loc1[l1]-loc2[l2]))
        if loc1[l1] < loc2[l2] {
            l1++
        } else {
            l2++
        }
    }
    return res
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
