func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    mapS := make(map[byte]byte)
    mapT := make(map[byte]byte)
    for i := 0; i < len(s); i++ {
        if v, ok := mapS[s[i]]; ok && v != t[i] {
            return false
        }
        if v, ok := mapT[t[i]]; ok && v != s[i] {
            return false
        }
        mapS[s[i]] = t[i]
        mapT[t[i]] = s[i]
    }
    return true
}
