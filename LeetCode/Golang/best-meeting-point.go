func minTotalDistance(grid [][]int) int {
    m := len(grid)
    if m == 0 {
        return 0
    }
    n := len(grid[0])
    if n == 0 {
        return 0
    }
    coorXs := make([]int, 0)
    coorYs := make([]int, 0)
    people := make([][2]int, 0)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                coorXs = append(coorXs, i)
                coorYs = append(coorYs, j)
                people = append(people, [2]int{i,j})
            }
        }
    }
    sort.Ints(coorXs)
    sort.Ints(coorYs)
    dist, meetX, meetY := 0, coorXs[len(coorXs)/2], coorYs[len(coorYs)/2]
    for _, person := range people {
        px, py := person[0], person[1]
        dist += abs(px - meetX) + abs(py - meetY)
    }
    return dist
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
