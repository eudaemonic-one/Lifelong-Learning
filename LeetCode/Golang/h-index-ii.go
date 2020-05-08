func hIndex(citations []int) int {
    N := len(citations)
    if N == 0 || (N == 1 && citations[0] < 1) {
        return 0
    }
    l, r := 0, N
    for l < r {
        m := (l+r) / 2
        if citations[m] == N-m {
            return N - m
        } else if citations[m] < N-m {
            l = m + 1
        } else {
            r = m
        }
    }
    return N - l
}
