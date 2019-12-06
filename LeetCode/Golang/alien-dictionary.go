func alienOrder(words []string) string {
    if len(words) == 0 {
        return ""
    }
    // 0=todo 1=doing 2=done
    visited := make(map[byte]int)
    edges := make(map[[2]byte]bool)
    order := ""
    buildGraph(words, visited, edges)
    for vertice, status := range visited {
        if status == 0 {
            if !dfs(vertice, visited, edges, &order) {
                return ""
            }
        }
    }
    return reverse(order)
}

func buildGraph(words []string, visited map[byte]int, edges map[[2]byte]bool) {
    for i, word := range words {
        for _, c := range word {
            visited[byte(c)] = 0
        }
        if i > 0 {
            for j := 0; j < min(len(words[i-1]), len(words[i])); j++ {
                if words[i-1][j] != words[i][j] {
                    edges[[2]byte{words[i-1][j],words[i][j]}] = true
                    break
                }
            }
        }
    }
}

func dfs(start byte, visited map[byte]int, edges map[[2]byte]bool, order *string) bool {
    visited[start] = 1
    for end := range visited {
        if edges[[2]byte{start,end}] == true {
            if visited[end] == 1 {
                return false
            } else if visited[end] == 0 {
                if !dfs(end, visited, edges, order) {
                    return false
                }
            }
        }
    }
    visited[start] = 2
    *order += string(start)
    return true
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func reverse(s string) string {
    runes := []rune(s)
    for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
        runes[l], runes[r] = runes[r], runes[l]
    }
    return string(runes)
}
