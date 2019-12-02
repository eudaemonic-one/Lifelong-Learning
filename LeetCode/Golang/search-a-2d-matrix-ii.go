func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    if m == 0 {
        return false
    }
    n := len(matrix[0])
    if n == 0 {
        return false
    }
    if target < matrix[0][0] || target > matrix[m-1][n-1] {
        return false
    }
    for i := 0; i < m; i++ {
        if target > matrix[i][n-1] {
            continue
        }
        l, r := 0, n-1
        for l <= r {
            m := (l+r) / 2
            if matrix[i][m] == target {
                return true
            } else if matrix[i][m] < target {
                l = m + 1
            } else {
                r = m - 1
            }
        }
    }
    return false
}
