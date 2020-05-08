func getFactors(n int) [][]int {
    res := make([][]int, 0)
    backtrack(n, []int{}, &res)
    return res
}

func backtrack(n int, factors []int, res *[][]int) {
    if n == 1 {
        if len(factors) > 1 {
            *res = append(*res, factors)
        }
        return
    }
    start := 2
    if len(factors) > 0 {
        start = factors[len(factors)-1]
    }
    for i := start; i <= n; i++ {
        if n % i == 0 {
            factors = append(factors, i)
            backtrack(n/i, factors, res)
            factors = append([]int{}, factors[:len(factors)-1]...)
        }
    }
}
