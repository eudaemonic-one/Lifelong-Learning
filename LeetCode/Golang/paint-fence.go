func numWays(n int, k int) int {
    if n == 0 {
        return 0
    } else if n == 1 {
        return k
    }
    same := k
    diff := k * (k-1)
    for i := 2; i < n; i++ {
        diff, same = (same + diff) * (k-1), diff
    }
    return same + diff
}
