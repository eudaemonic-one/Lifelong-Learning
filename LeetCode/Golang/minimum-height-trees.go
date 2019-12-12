func findMinHeightTrees(n int, edges [][]int) []int {
    // corner case
    if n == 0 {
        return []int{}
    }
    if n == 1 {
        return []int{0}
    }
    // initialize the undirected graph
    graph := make([]map[int]int, n)
    for i := 0; i < n; i++ {
        graph[i] = make(map[int]int)
    }
    for _, edge := range edges {
        s, e := edge[0], edge[1]
        graph[s][e] = e
        graph[e][s] = s
    }
    // find leaves layer
    currLayer := make([]int, 0)
    for i := 0; i < len(graph); i++ {
        if len(graph[i]) == 1 {
            currLayer = append(currLayer, i)
        }
    }
    // bfs
    for true {
        nextLayer := make([]int, 0)
        for _, leaf := range currLayer {
            for _, nei := range graph[leaf] {
                delete(graph[nei], leaf)
                if len(graph[nei]) == 1 {
                    nextLayer = append(nextLayer, nei)
                }
            }
        }
        if len(nextLayer) == 0 {
            return currLayer
        }
        currLayer = nextLayer
    }
    return currLayer
}
