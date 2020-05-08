const N int = 26

type TrieNode struct {
    Children [N]*TrieNode
    Word string
}

func (p *TrieNode) search() {
    
}

func findWords(board [][]byte, words []string) []string {
    res := make([]string, 0)
    root := buildTrie(words)
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            dfs(board, i, j, root, &res)
        }
    }
    return res
}

func dfs(board [][]byte, x, y int, node *TrieNode, res *[]string) {
    if (x < 0 || x >= len(board) ||
        y < 0 || y >= len(board[0])) {
        return
    }
    c := board[x][y]
    i := int(c) - int('a')
    if c == '#' || node.Children[i] == nil {
        return
    }
    node = node.Children[i]
    if node.Word != "" {
        *res = append(*res, node.Word)
        node.Word = ""
    }
    board[x][y] = '#'
    dfs(board, x-1, y, node, res)
    dfs(board, x, y-1, node, res)
    dfs(board, x, y+1, node, res)
    dfs(board, x+1, y, node, res)
    board[x][y] = c
    return
}

func buildTrie(words []string) *TrieNode {
    root := &TrieNode{}
    for _, word := range words {
        cur := root
        for _, c := range word {
            i := int(c) - int('a')
            if cur.Children[i] == nil {
                cur.Children[i] = &TrieNode{}
            }
            cur = cur.Children[i]
        }
        cur.Word = word
    }
    return root
}
