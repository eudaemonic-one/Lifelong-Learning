func minTimeToVisitAllPoints(points [][]int) int {
    if len(points) == 0 {
        return 0
    }
    res := 0
    x, y := points[0][0], points[0][1]
    for i := 1; i < len(points); i++ {
        nx, ny := points[i][0], points[i][1]
        dx, dy := nx - x, ny - y
        if dx < 0 {
            dx = -dx
        }
        if dy < 0 {
            dy = -dy
        }
        if dx > dy {
            res += dx
        } else {
            res += dy
        }
        x, y = nx, ny
    }
    return res
}
