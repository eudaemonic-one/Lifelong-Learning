func wordPatternMatch(pattern string, str string) bool {
    patDict := make(map[byte]string)
    strDict := make(map[string]byte)
    return backtrack(pattern, str, patDict, strDict, 0, 0)
}

func backtrack(pattern, str string, patDict map[byte]string, strDict map[string]byte, p, s int) bool {
    if p == len(pattern) && s == len(str) {
        return true
    }
    if p == len(pattern) || s == len(str) {
        return false
    }
    if word, ok := patDict[pattern[p]]; !ok {
        for e := s+1; e <= len(str); e++ {
            candidate := str[s:e]
            if _, ok2 := strDict[candidate]; ok2 {
                continue
            }
            patDict[pattern[p]] = candidate
            strDict[candidate] = pattern[p]
            if backtrack(pattern, str, patDict, strDict, p+1, e) {
                return true
            }
            delete(patDict, pattern[p])
            delete(strDict, candidate)
        }
    } else { // ok
        e := s+len(word)
        if e <= len(str) && word != str[s:e] {
            return false
        }
        return backtrack(pattern, str, patDict, strDict, p+1, e)
    }
    return false
}
