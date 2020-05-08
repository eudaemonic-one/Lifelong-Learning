const INF = 2147483647

func wallsAndGates(rooms [][]int)  {
    m := len(rooms)
    if m == 0 {
        return
    }
    n := len(rooms[0])
    if n == 0 {
        return
    }
    queue := make([][2]int, 0)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if rooms[i][j] == 0 {
                queue = append(queue, [2]int{i,j})
            }
        }
    }
    for len(queue) > 0 {
        gate := queue[0]
        queue = queue[1:]
        gx, gy := gate[0], gate[1]
        for _, d := range [][2]int{{-1,0},{0,-1},{0,1},{1,0}} {
            nx, ny := gx + d[0], gy + d[1]
            if 0 <= nx && nx < m && 0 <= ny && ny < n && rooms[nx][ny] == INF {
                rooms[nx][ny] = rooms[gx][gy] + 1
                queue = append(queue, [2]int{nx,ny})
            }
        }
    }
}
