type WordDistance struct {
    distMap map[string][]int
}


func Constructor(words []string) WordDistance {
    dist := WordDistance{make(map[string][]int)}
    for i, word := range words {
        dist.distMap[word] = append(dist.distMap[word], i)
    }
    return dist
}


func (this *WordDistance) Shortest(word1 string, word2 string) int {
    loc1, loc2 := this.distMap[word1], this.distMap[word2]
    l1, l2 := 0, 0
    minDiff := int(^uint(0) >> 1)
    for l1 < len(loc1) && l2 < len(loc2) {
        minDiff = min(minDiff, abs(loc1[l1]-loc2[l2]))
        if loc1[l1] < loc2[l2] {
            l1 += 1
        } else {
            l2 += 1
        }
    }
    return minDiff
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

/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Shortest(word1,word2);
 */
