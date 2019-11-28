func rangeBitwiseAnd(m int, n int) int {
    var i, j int
    for int(math.Pow(2, float64(i))) <= m {
        i += 1
    }
    for int(math.Pow(2, float64(j))) <= n {
        j += 1
    }
    if i != j {
        return 0
    }
    res := m
    for x := m+1; x <= n; x++ {
        res &= x
    }
    return res
}
