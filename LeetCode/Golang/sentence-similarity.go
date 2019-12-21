func areSentencesSimilar(words1 []string, words2 []string, pairs [][]string) bool {
    if len(words1) != len(words2) {
        return false
    }
    dict := make(map[string][]string)
    for _, pair := range pairs {
        w1, w2 := pair[0], pair[1]
        dict[w1] = append(dict[w1], w2)
        dict[w2] = append(dict[w2], w1)
    }
    for i := 0; i < len(words2); i++ {
        w1, w2 := words1[i], words2[i]
        if w1 != w2 {
            flag := false
            for _, w := range dict[w1] {
                if w == w2 {
                    flag = true
                    break
                }
            }
            if flag {
                continue
            }
            for _, w := range dict[w2] {
                if w == w1 {
                    flag = true
                    break
                }
            }
            if !flag {
                return false
            }
        }
    }
    return true
}
