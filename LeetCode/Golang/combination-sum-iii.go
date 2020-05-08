func combinationSum3(k int, n int) [][]int {
    res := make([][]int, 0)
    backtrack(k, n, 1, []int{}, &res)
    return res
}

func backtrack(k, n, idx int, path []int, res *[][]int) {
    if k == 0 || n == 0 || idx == 10 {
        if k == 0 && n == 0 {
            tmp := make([]int, len(path))
            copy(tmp, path)
            *res = append(*res, tmp)
        }
        return
    }
    for i := idx; i <= min(9, n); i++ {
        path = append(path, i)
        backtrack(k-1, n-i, i+1, path, res)
        path = path[:len(path)-1]
    }
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
