func gameOfLife(board [][]int)  {
    m := len(board)
    if m == 0 {
        return
    }
    n := len(board[0])
    if n == 0 {
        return
    }
    neis := [8][2]int{{-1,0},{-1,1},{0,1},{1,1},{1,0},{1,-1},{0,-1},{-1,-1}}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            neighbors := 0
            for _, d := range neis {
                nx, ny := i + d[0], j + d[1]
                if 0 <= nx && nx < m && 0 <= ny && ny < n && (board[nx][ny] == 1 || board[nx][ny] == -1) {
                    neighbors++
                }
            }
            if board[i][j] == 1 && (neighbors < 2 || neighbors > 3) {
                board[i][j] = -1 // Rule 1 & 3 Under-population or over-population
            } else if board[i][j] == 0 && neighbors == 3 {
                board[i][j] = 2 // Rule 4 Reproduction
            }
        }
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] > 0 {
                board[i][j] = 1
            } else {
                board[i][j] = 0
            }
        }
    }
}
