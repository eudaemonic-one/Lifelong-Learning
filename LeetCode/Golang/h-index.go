func hIndex(citations []int) int {
    N := len(citations)
    counts := make([]int, N+1)
    for _, n := range citations {
        counts[min(n, N)]++
    }
    n := N
    for sum := counts[n]; n > sum; sum += counts[n] {
        n--
    }
    return n
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
