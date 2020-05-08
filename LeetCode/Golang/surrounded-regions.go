func solve(board [][]byte)  {
    if (len(board) == 0 || len(board[0]) == 0) {
        return
    }
    for i := 0; i < len(board); i++ {
        dfs(board, 'E', i, 0)
        dfs(board, 'E', i, len(board[0])-1)
    }
    for j := 0; j < len(board[0]); j++ {
        dfs(board, 'E', 0, j)
        dfs(board, 'E', len(board)-1, j)
    }
    for i := 1; i < len(board)-1; i++ {
        for j := 1; j < len(board[0])-1; j++ {
            dfs(board, 'X', i, j)
        }
    }
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if (board[i][j] == 'E') {
                board[i][j] = 'O'
            }
        }
    }
    return
}

func dfs(board [][]byte, sign byte, x int, y int) {
    var flag = false
    if (board[x][y] == 'O') {
        board[x][y] = sign
        flag = true
    }
    for _, d := range [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
        if (flag) {
            nx := x+d[0]
            ny := y+d[1]
            if (0 <= nx && nx < len(board) && 0 <= ny && ny < len(board[0])) {
                dfs(board, sign, nx, ny)
            }
        }
    }
}
