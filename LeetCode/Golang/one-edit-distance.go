func isOneEditDistance(s string, t string) bool {
    lenDiff := len(s) - len(t)
    if lenDiff > 1 || lenDiff < -1 || s == t {
        return false
    }
    if lenDiff > 0 {
        s, t = t, s
    }
    if lenDiff == 0 {
        dist := 0
        for i := 0; i < len(s); i++ {
            if s[i] != t[i] {
                dist++
            }
            if dist > 1 {
                return false
            }
        }
    } else {
        dist := 0
        for i, j := 0, 0; i < len(s) && j < len(t); {
            if s[i] == t[j] {
                i++
                j++
            } else {
                dist++
                j++
            }
            if dist > 1 {
                return false
            }
        }
    }
    return true
}
