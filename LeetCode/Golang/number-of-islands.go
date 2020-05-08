func numIslands(grid [][]byte) int {
    var res int
    m := len(grid)
    if m == 0 {
        return 0
    }
    n := len(grid[0])
    if n == 0 {
        return 0
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' {
                dfs(grid, i, j, m, n)
                res += 1
            }
        }
    }
    return res
}

func dfs(grid [][]byte, x, y, m, n int) {
    grid[x][y] = '0'
    for _, d := range [4][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}} {
        nx, ny := x + d[0], y + d[1]
        if 0 <= nx && nx < m && 0 <= ny && ny < n && grid[nx][ny] == '1' {
            dfs(grid, nx, ny, m, n)
        }
    }
    return
}
