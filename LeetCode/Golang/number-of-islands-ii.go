func numIslands2(m int, n int, positions [][]int) []int {
    islands := make([]int, 0)
    parents := make([]int, m*n)
    dx := [4]int{-1,0,0,1}
    dy := [4]int{0,1,-1,0}
    for i := 0; i < m*n; i++ {
        parents[i] = -1
    }
    for i := 0; i < len(positions); i++ {
        count := 1
        if i > 0 {
            count = islands[len(islands)-1] + 1
        }
        px, py := positions[i][0], positions[i][1]
        if parents[px*n+py] != -1 {
            islands = append(islands, count-1)
            continue
        }
        parents[px*n+py] = px*n+py
        for j := 0; j < 4; j++ {
            nx := px + dx[j]
            ny := py + dy[j]
            if 0 <= nx && nx < m && 0 <= ny && ny < n && parents[nx*n+ny] != -1 {
                root1 := find(parents, nx*n+ny)
                root2 := parents[px*n+py]
                if root1 != root2 { // 2 isolated islands
                    count--
                }
                parents[root1] = root2 // union
            }
        }
        islands = append(islands, count)
    }
    return islands
}

func find(nums []int, num int) int {
    if nums[num] == num {
        return num
    }
    return find(nums, nums[num])
}
