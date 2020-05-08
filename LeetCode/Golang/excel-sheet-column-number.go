func titleToNumber(s string) int {
    var res int
    var base int = 1
    for i := len(s)-1; i >= 0; i-- {
        res += (int(s[i] - 'A') + 1) * base
        base *= 26
    }
    return res
}
