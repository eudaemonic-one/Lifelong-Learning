func canFinish(numCourses int, prerequisites [][]int) bool {
    // Edge cases
    if numCourses <= 1 {
        return true
    }
    // Build graph
    edges := make(map[int]([]int), 0)
    for _, pair := range prerequisites {
        if _, ok := edges[pair[1]]; !ok {
            edges[pair[1]] = []int{pair[0]}
        } else {
            edges[pair[1]] = append(edges[pair[1]], pair[0])
        }
    }
    // Initiate auxiliaries
    todo := make([]bool, numCourses)
    done := make([]bool, numCourses)
    // Judge if there is cycle in the graph
    for i := 0; i < numCourses; i++ {
        if _, ok := edges[i]; ok && !done[i] && isCyclic(edges, todo, done, i) {
            return false
        }
    }
    return true
}

func isCyclic(edges map[int][]int, todo []bool, done []bool, v int) bool {
    if done[v] {
        return true
    }
    done[v] = true
    for _, e := range edges[v] {
        if isCyclic(edges, todo, done, e) {
            return true
        }
    }
    done[v] = false
    return false
}
