func trailingZeroes(n int) int {
    if n == 0 {
        return 0
    }
    var cnt int
    for n > 0 {
        cnt += n / 5
        n = n / 5
    }
    return cnt
}
