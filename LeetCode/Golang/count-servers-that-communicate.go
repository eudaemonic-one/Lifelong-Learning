func countServers(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    res := 0
    m, n := len(grid), len(grid[0])
    rows := make([]int, m)
    cols := make([]int, n)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                rows[i] += 1
                cols[j] += 1
            }
        }
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 && (rows[i] > 1 || cols[j] > 1) {
                res += 1
            }
        }
    }
    return res
}
