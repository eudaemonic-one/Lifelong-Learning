func exist(board [][]byte, word string) bool {
    if (len(board) == 0 ||
        len(board[0]) == 0 ||
        len(word) == 0) {
        // Edge cases
        return false
    }
    visited := make([][]bool, len(board))
    for i := 0; i < len(board); i++ {
        visited[i] = make([]bool, len(board[0]))
    }
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if board[i][j] == word[0] && backtrack(board, visited, i, j, word){
                return true
            }
        }
    }
    return false
}

func backtrack(board [][]byte, visited [][]bool, x int, y int, word string) bool {
    if len(word) == 0 {
        return true
    }
    if (x < 0 || x >= len(board) || 
        y < 0 || y >= len(board[0]) || 
        visited[x][y] || 
        board[x][y] != word[0]) {
        // Word search fails for coordinate (x, y)
        return false
    }
    visited[x][y] = true
    if (backtrack(board, visited, x-1, y, word[1:]) || 
        backtrack(board, visited, x, y-1, word[1:]) || 
        backtrack(board, visited, x, y+1, word[1:]) || 
        backtrack(board, visited, x+1, y, word[1:])) {
        // Condition moves to adjacent cells
        return true
    }
    visited[x][y] = false
    return false
}
